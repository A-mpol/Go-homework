package converter

import (
	serviceModel "order/internal/model"
	order_v1 "shared/pkg/proto/order/v1"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func CreateRequestToModel(cancelRequest *order_v1.CreateRequest) serviceModel.CreateRequest {
	return serviceModel.CreateRequest{
		UserUuid:  cancelRequest.UserUuid,
		PartUuids: cancelRequest.PartUuids,
	}
}

func ServiceCreateResponseToApi(serviceModelCreateResponse serviceModel.CreateResponse) *order_v1.CreateResponse {
	return &order_v1.CreateResponse{
		OrderUuid:  serviceModelCreateResponse.OrderUuid,
		TotalPrice: serviceModelCreateResponse.TotalPrice,
	}
}

func GetRequestToModel(getRequest *order_v1.GetRequest) serviceModel.GetRequest {
	return serviceModel.GetRequest{
		OrderUuid: getRequest.OrderUuid,
	}
}

func ServiceGetResponseToApi(serviceModelGetResponse serviceModel.GetResponse) *order_v1.GetResponse {
	return &order_v1.GetResponse{
		Order: ServiceOrderToApi(serviceModelGetResponse.Order),
	}
}

func ServiceOrderToApi(serviceModelOrder serviceModel.Order) *order_v1.Order {
	return &order_v1.Order{
		OrderUuid:  serviceModelOrder.OrderUuid,
		UserUuid:   serviceModelOrder.UserUuid,
		PartUuids:  serviceModelOrder.PartUuids,
		TotalPrice: serviceModelOrder.TotalPrice,
		TransactionUuid: &wrapperspb.StringValue{
			Value: serviceModelOrder.TransactionUuid,
		},
		PaymentMethod: order_v1.PaymentMethod(serviceModelOrder.PaymentMethod),
		Status:        order_v1.Status(serviceModelOrder.Status),
	}
}

func PayRequestToModel(payRequest *order_v1.PayRequest) serviceModel.PayRequest {
	return serviceModel.PayRequest{
		OrderUuid:     payRequest.OrderUuid,
		PaymentMethod: serviceModel.PaymentMethod(payRequest.PaymentMethod),
	}
}

func ServicePayResponseToApi(serviceModelPayResponse serviceModel.PayResponse) *order_v1.PayResponse {
	return &order_v1.PayResponse{
		TransactionUuid: serviceModelPayResponse.TransactionUuid,
	}
}

func CancelRequestToModel(cancelRequest *order_v1.CancelRequest) serviceModel.CancelRequest {
	return serviceModel.CancelRequest{
		OrderUuid: cancelRequest.OrderUuid,
	}
}
