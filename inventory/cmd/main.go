package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	inventory_v1 "shared/pkg/proto/inventory/v1"
	"sync"
	"syscall"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const grpcPort = 50052

type inventoryService struct {
	inventory_v1.UnimplementedInventoryServiceServer

	mu        sync.RWMutex
	inventory map[string]*inventory_v1.Part
}

func (i *inventoryService) GetPart(ctx context.Context, in *inventory_v1.GetPartRequest) (*inventory_v1.GetPartResponse, error) {
	i.mu.RLock()
	part, ok := i.inventory[in.Uuid]
	i.mu.RUnlock()
	if !ok {
		return nil, status.Errorf(codes.NotFound, "part with UUID %s not found", in.GetUuid())
	}

	return &inventory_v1.GetPartResponse{
		Part: part,
	}, nil
}

func (inv *inventoryService) ListParts(ctx context.Context, in *inventory_v1.ListPartsRequest) (*inventory_v1.ListPartsResponse, error) {
	listParts := []*inventory_v1.Part{}

	inv.mu.RLock()
	for uuid := range inv.inventory {
		fitsParameter := true
		for i, checkUuid := range in.Filter.Uuids {
			if uuid == checkUuid {
				break
			}
			if uuid != checkUuid && i == len(in.Filter.Uuids)-1 {
				fitsParameter = false
			}
		}

		for i, checkName := range in.Filter.Names {
			if !fitsParameter || inv.inventory[uuid].Name == checkName {
				break
			}
			if inv.inventory[uuid].Name != checkName && i == len(in.Filter.Names)-1 {
				fitsParameter = false
			}
		}

		for i, checkCategory := range in.Filter.Categories {
			if !fitsParameter || inv.inventory[uuid].Category == checkCategory {
				break
			}
			if inv.inventory[uuid].Category != checkCategory && i == len(in.Filter.Categories)-1 {
				fitsParameter = false
			}
		}

		for i, checkManufacturerCountrie := range in.Filter.ManufacturerCountries {
			if !fitsParameter || inv.inventory[uuid].Manufacturer.Country == checkManufacturerCountrie {
				break
			}
			if inv.inventory[uuid].Manufacturer.Country != checkManufacturerCountrie && i == len(in.Filter.ManufacturerCountries)-1 {
				fitsParameter = false
			}
		}

		for i, checkTag := range in.Filter.Tags {
			if !fitsParameter {
				break
			}

			noSuchTag := true
			for _, tag := range inv.inventory[uuid].Tags {
				if checkTag == tag {
					noSuchTag = false
					break
				}
			}
			if i == len(in.Filter.ManufacturerCountries)-1 && noSuchTag {
				fitsParameter = false
			}
		}

		if fitsParameter {
			listParts = append(listParts, inv.inventory[uuid])
		}
	}
	inv.mu.RUnlock()

	return &inventory_v1.ListPartsResponse{
		Parts: listParts,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()

	parts := make([]*inventory_v1.Part, 0)
	for range 100 {
		parts = append(parts, &inventory_v1.Part{
			Uuid:          uuid.NewString(),
			Name:          gofakeit.Name(),
			Description:   gofakeit.Letter(),
			Price:         10.0,
			StockQuantity: 10.0,
			Category:      inventory_v1.Category_CATEGORY_ENGINE,
			Dimensions: &inventory_v1.Dimensions{
				Length: 10,
				Height: 10,
				Width:  6,
				Weight: 7,
			},
		})
	}

	m := make(map[string]*inventory_v1.Part)
	for _, p := range parts {
		m[p.GetUuid()] = p
	}
	service := &inventoryService{
		inventory: m,
	}

	inventory_v1.RegisterInventoryServiceServer(s, service)

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
