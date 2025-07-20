package converter

import (
	serviceModel "order/internal/model"
	repoModel "order/internal/repository/model"
)

func ServiceCreateRequestToRepo(serviceModelCreateRequest serviceModel.CreateRequest) repoModel.CreateRequest {
	return repoModel.CreateRequest{
		UserUuid:  serviceModelCreateRequest.UserUuid,
		PartUuids: serviceModelCreateRequest.PartUuids,
	}
}

func RepoCreateResponseToService(repoModelCreateResponse repoModel.CreateResponse) serviceModel.CreateResponse {
	return serviceModel.CreateResponse{
		OrderUuid:  repoModelCreateResponse.OrderUuid,
		TotalPrice: repoModelCreateResponse.TotalPrice,
	}
}

func ServiceGetRequestToRepo(serviceModelGetRequest serviceModel.GetRequest) repoModel.GetRequest {
	return repoModel.GetRequest{
		OrderUuid: serviceModelGetRequest.OrderUuid,
	}
}

func RepoGetResponseToService(repoModelGetResponse repoModel.GetResponse) serviceModel.GetResponse {
	return serviceModel.GetResponse{
		Order: RepoOrderToService(repoModelGetResponse.Order),
	}
}

func RepoOrderToService(repoModelOrder repoModel.Order) serviceModel.Order {
	return serviceModel.Order{
		OrderUuid:       repoModelOrder.OrderUuid,
		UserUuid:        repoModelOrder.UserUuid,
		PartUuids:       repoModelOrder.PartUuids,
		TotalPrice:      repoModelOrder.TotalPrice,
		TransactionUuid: repoModelOrder.TransactionUuid,
		PaymentMethod:   serviceModel.PaymentMethod(repoModelOrder.PaymentMethod),
		Status:          serviceModel.Status(repoModelOrder.Status),
	}
}

func ServiceUpdateRequestToRepo(serviceModelUpdateRequest serviceModel.UpdateRequest) repoModel.UpdateRequest {
	return repoModel.UpdateRequest{
		OrderUuid:       serviceModelUpdateRequest.OrderUuid,
		TransactionUuid: serviceModelUpdateRequest.TransactionUuid,
		PaymentMethod:   repoModel.PaymentMethod(serviceModelUpdateRequest.PaymentMethod),
		Status:          repoModel.Status(serviceModelUpdateRequest.Status),
	}
}
