package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	inventoryApi "inventory/internal/api/inventory/v1"
	inventoryRepository "inventory/internal/repository/inventory"
	inventoryService "inventory/internal/service/inventory"
)

const grpcPort = 50051

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()

	repo := inventoryRepository.NewRepository()
	service := inventoryService.NewService(repo)
	api := inventoryApi.NewApi(service)

	inventory_v1.RegisterInventoryServiceServer(s, api)

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
