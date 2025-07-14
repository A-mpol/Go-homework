package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	payment_v1 "shared/pkg/proto/payment/v1"
	"syscall"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50053

type PaymentService struct {
	payment_v1.UnimplementedPaymentServiceServer
}

func (p *PaymentService) PayOrder(ctx context.Context, in *payment_v1.PayOrderRequest) (*payment_v1.PayOrderResponse, error) {
	uuid := uuid.NewString()
	log.Printf("Оплата прошла успешно, transaction_uuid: %s", uuid)
	return &payment_v1.PayOrderResponse{
		TransactionUuid: uuid,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()

	service := &PaymentService{}

	payment_v1.RegisterPaymentServiceServer(s, service)

	reflection.Register(s)

	go func() {
		log.Printf("🚀 gRPC server listening on %d\n", grpcPort)
		err = s.Serve(lis)
		if err != nil {
			log.Printf("failed to serve: %v\n", err)
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down gRPC server...")
	s.GracefulStop()
	lis.Close()
	log.Println("✅ Server stopped")
}
