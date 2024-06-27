package dtos

type ReceiveOrderDTO struct {
	MerchantID string `json:"merchant_id"`
	OrderID    string `json:"order_id"`
	Status     string `json:"status"`
}
