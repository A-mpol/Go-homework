package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	order_v1 "shared/pkg/proto/order/v1"
	payment_v1 "shared/pkg/proto/payment/v1"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	grpcPort               = 50052
	httpPort               = 8080
	inventoryServerAddress = "localhost:50051"
	paymentServerAddress   = "localhost:50053"
	orderServerAddress     = "localhost:50054"
)

type orderService struct {
	order_v1.UnimplementedOrderServiceServer

	mu     sync.RWMutex
	orders map[string]*order_v1.Order

	inventoryClient inventory_v1.InventoryServiceClient
	paymentClient   payment_v1.PaymentServiceClient
}

func NewOrderService(
	inventoryClient inventory_v1.InventoryServiceClient,
	paymentClient payment_v1.PaymentServiceClient,
) *orderService {
	return &orderService{
		orders:          make(map[string]*order_v1.Order),
		inventoryClient: inventoryClient,
		paymentClient:   paymentClient,
	}
}

func (o *orderService) Create(ctx context.Context, in *order_v1.CreateRequest) (*order_v1.CreateResponse, error) {
	listParts, err := o.inventoryClient.ListParts(ctx, &inventory_v1.ListPartsRequest{
		Filter: &inventory_v1.PartsFilter{
			Uuids: in.GetPartUuids(),
		}})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Don't get list parts")
	}
	if len(listParts.Parts) < len(in.GetPartUuids()) {
		return nil, status.Errorf(codes.Internal, "Not all details exist")
	}

	var totalPrice float64
	for _, part := range listParts.Parts {
		totalPrice += part.Price
	}

	orderUuid := uuid.NewString()

	o.mu.Lock()
	o.orders[orderUuid] = &order_v1.Order{
		OrderUuid:  orderUuid,
		UserUuid:   in.GetUserUuid(),
		PartUuids:  in.GetPartUuids(),
		TotalPrice: totalPrice,
		Status:     order_v1.Status_STATUS_PENDING_PAYMENT,
	}
	o.mu.Unlock()

	return &order_v1.CreateResponse{
		OrderUuid:  orderUuid,
		TotalPrice: totalPrice,
	}, nil
}

func (o *orderService) Pay(ctx context.Context, in *order_v1.PayRequest) (*order_v1.PayResponse, error) {
	o.mu.RLock()
	order, ok := o.orders[in.GetOrderUuid()]
	o.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	paymentInformation, err := o.paymentClient.PayOrder(ctx, &payment_v1.PayOrderRequest{
		OrderUuid:     in.GetOrderUuid(),
		UserUuid:      order.UserUuid,
		PaymentMethod: payment_v1.PaymentMethod(in.GetPaymentMethod()),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to pay")
	}

	o.mu.Lock()
	o.orders[in.GetOrderUuid()].TransactionUuid = &wrapperspb.StringValue{Value: paymentInformation.TransactionUuid}
	o.orders[in.GetOrderUuid()].PaymentMethod = in.GetPaymentMethod()
	o.orders[in.GetOrderUuid()].Status = order_v1.Status_STATUS_PAID
	o.mu.Unlock()

	return &order_v1.PayResponse{
		TransactionUuid: paymentInformation.TransactionUuid,
	}, nil
}

func (o *orderService) Get(ctx context.Context, in *order_v1.GetRequest) (*order_v1.GetResponse, error) {
	o.mu.RLock()
	order, ok := o.orders[in.GetOrderUuid()]
	o.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	return &order_v1.GetResponse{
		Order: order,
	}, nil
}

func (o *orderService) Cancel(ctx context.Context, in *order_v1.CancelRequest) (*emptypb.Empty, error) {
	o.mu.RLock()
	order, ok := o.orders[in.GetOrderUuid()]
	o.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Order not found")
	}

	if order.Status == order_v1.Status_STATUS_PAID {
		return nil, status.Errorf(codes.AlreadyExists, "Order already paid, cannot cancel")
	}

	o.mu.Lock()
	o.orders[in.GetOrderUuid()].Status = order_v1.Status_STATUS_CANCELLED
	o.mu.Unlock()

	return &emptypb.Empty{}, nil
}

func main() {
	// Запускаем gRPC сервер
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}
	defer func() {
		if cerr := lis.Close(); cerr != nil {
			log.Printf("failed to close listener: %v\n", cerr)
		}
	}()

	inventoryConn, err := grpc.NewClient(
		inventoryServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	paymentConn, err := grpc.NewClient(
		paymentServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	orderConn, err := grpc.NewClient(
		orderServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	_ = order_v1.NewOrderServiceClient(orderConn)

	service := NewOrderService(
		inventory_v1.NewInventoryServiceClient(inventoryConn),
		payment_v1.NewPaymentServiceClient(paymentConn),
	)

	s := grpc.NewServer()
	order_v1.RegisterOrderServiceServer(s, service)
	reflection.Register(s)

	// Запускаем gRPC сервер в горутине
	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	// Запускаем HTTP сервер с gRPC Gateway и Swagger UI
	var gwServer *http.Server
	go func() {
		// Создаем контекст с отменой
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Создаем мультиплексор для HTTP запросов
		mux := runtime.NewServeMux()

		// Настраиваем опции для соединения с gRPC сервером
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

		// Регистрируем gRPC-gateway хендлеры
		err = order_v1.RegisterOrderServiceHandlerFromEndpoint(
			ctx,
			mux,
			fmt.Sprintf("localhost:%d", grpcPort),
			opts,
		)
		if err != nil {
			log.Printf("Failed to register gateway: %v\n", err)
			return
		}

		// Создаем файловый сервер для swagger-ui
		fileServer := http.FileServer(http.Dir("../../shared/api"))

		// Создаем HTTP маршрутизатор
		httpMux := http.NewServeMux()

		// Регистрируем API эндпоинты
		httpMux.Handle("/api/", mux)

		// Swagger UI эндпоинты
		httpMux.Handle("/swagger-ui.html", fileServer)
		httpMux.Handle("/swagger.json", fileServer)

		// Редирект с корня на Swagger UI
		httpMux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				http.Redirect(w, r, "/swagger-ui.html", http.StatusMovedPermanently)
				return
			}
			fileServer.ServeHTTP(w, r)
		}))

		// Создаем HTTP сервер
		gwServer = &http.Server{
			Addr:              fmt.Sprintf(":%d", httpPort),
			Handler:           httpMux,
			ReadHeaderTimeout: 10 * time.Second,
		}

		// Запускаем HTTP сервер
		log.Printf("🌐 HTTP server with gRPC-Gateway and Swagger UI listening on %d\n", httpPort)
		err = gwServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Failed to serve HTTP: %v\n", err)
			return
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down servers...")

	// Сначала аккуратно останавливаем HTTP сервер
	if gwServer != nil {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := gwServer.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
		log.Println("✅ HTTP server stopped")
	}

	// В конце останавливаем gRPC сервер
	s.GracefulStop()
	log.Println("✅ gRPC server stopped")
}
