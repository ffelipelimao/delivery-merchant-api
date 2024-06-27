package receiveorder

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/repository"
)

type ReceiveOrderUseCase struct {
	merchantOrderRepository repository.MerchantOrderRepository
}

func NewReceiveOrderUseCase(merchantOrderRepository repository.MerchantOrderRepository) *ReceiveOrderUseCase {
	return &ReceiveOrderUseCase{
		merchantOrderRepository: merchantOrderRepository,
	}
}

func (r *ReceiveOrderUseCase) Handle(ctx context.Context, receiveOrder *dtos.ReceiveOrderDTO) error {
	merchantOrder := entities.NewMerchantOrder(receiveOrder.MerchantID, receiveOrder.Status, receiveOrder.OrderID)

	return r.merchantOrderRepository.Insert(ctx, merchantOrder)
}
