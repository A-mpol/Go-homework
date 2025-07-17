package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	paymentApi "payment/internal/api/payment/v1"
	paymentService "payment/internal/service/payment"
	payment_v1 "shared/pkg/proto/payment/v1"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50053

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()

	service := paymentService.NewService()
	api := paymentApi.NewApi(service)

	payment_v1.RegisterPaymentServiceServer(s, api)

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
