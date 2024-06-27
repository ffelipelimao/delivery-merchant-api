package entities

import (
	"errors"
	"time"
)

type MerchantOrder struct {
	MerchantID       string    `json:"merchant_id"`
	OrderID          string    `json:"order_id"`
	Status           string    `json:"status"`
	LastModifiedDate time.Time `json:"last_modified_date"`
	CreatedDate      time.Time `json:"created_date"`
}

func NewMerchantOrder(merchantID, status string, orderID string) *MerchantOrder {
	return &MerchantOrder{
		MerchantID: merchantID,
		OrderID:    orderID,
		Status:     status,
	}
}

func (m *MerchantOrder) ValidateStatus(statusReceive string) error {
	/*
		if _, ok := validateStatus[statusReceive]; !ok{
			return errors.New("invalid status")
		}
	*/
	if m.Status == "Finalizado" {
		// TODO handle this right
		return errors.New("order is already finish")
	}
	return nil
}
