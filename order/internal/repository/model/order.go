package model

type CreateRequest struct {
	UserUuid  string
	PartUuids []string
}

type CreateResponse struct {
	OrderUuid  string
	TotalPrice float64
}

type PayRequest struct {
	OrderUuid     string
	PaymentMethod PaymentMethod
}

type PaymentMethod int32

const (
	PaymentMethod_PAYMENT_METHOD_UNKNOWN_UNSPECIFIED PaymentMethod = 0
	PaymentMethod_PAYMENT_METHOD_CARD                PaymentMethod = 1
	PaymentMethod_PAYMENT_METHOD_SBP                 PaymentMethod = 2
	PaymentMethod_PAYMENT_METHOD_CREDIT_CARD         PaymentMethod = 3
	PaymentMethod_PAYMENT_METHOD_INVESTOR_MONEY      PaymentMethod = 4
)

type PayResponse struct {
	TransactionUuid string
}

type GetRequest struct {
	OrderUuid string
}

type GetResponse struct {
	Order Order
}

type Order struct {
	OrderUuid       string
	UserUuid        string
	PartUuids       []string
	TotalPrice      float64
	TransactionUuid string
	PaymentMethod   PaymentMethod
	Status          Status
}

type Status int32

const (
	Status_STATUS_UNKNOWN_UNSPECIFIED Status = 0
	Status_STATUS_PENDING_PAYMENT     Status = 1
	Status_STATUS_PAID                Status = 2
	Status_STATUS_CANCELLED           Status = 3
)

type CancelRequest struct {
	OrderUuid string
}

type UpdateRequest struct {
	OrderUuid       string
	TransactionUuid string
	PaymentMethod   PaymentMethod
	Status          Status
}
