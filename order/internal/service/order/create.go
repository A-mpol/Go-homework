package order

import (
	"context"
	serviceModel "order/internal/model"

	inventory_v1 "shared/pkg/proto/inventory/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Create(ctx context.Context, in serviceModel.CreateRequest) (serviceModel.CreateResponse, error) {
	listParts, err := s.inventoryClient.ListParts(ctx, &inventory_v1.ListPartsRequest{
		Filter: &inventory_v1.PartsFilter{
			Uuids: in.PartUuids,
		}})
	if err != nil {
		return serviceModel.CreateResponse{}, status.Errorf(codes.Internal, "Don't get list parts")
	}
	if len(listParts.Parts) < len(in.PartUuids) {
		return serviceModel.CreateResponse{}, status.Errorf(codes.Internal, "Not all details exist")
	}

	var totalPrice float64
	for _, part := range listParts.Parts {
		totalPrice += part.Price
	}

	createResponse, err := s.orderRepository.Create(ctx, in, totalPrice)
	if err != nil {
		return serviceModel.CreateResponse{}, err
	}

	return createResponse, nil
}
