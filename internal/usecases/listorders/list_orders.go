package listorders

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/repository"
)

type ListOrdersUseCase struct {
	merchantOrderRepository repository.MerchantOrderRepository
}

func NewListOrdersUseCase(merchantOrderRepository repository.MerchantOrderRepository) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		merchantOrderRepository: merchantOrderRepository,
	}
}

func (o *ListOrdersUseCase) Handle(ctx context.Context, merchantID string) (*[]entities.MerchantOrder, error) {
	statusReceive := "Pago"
	merchantOrder, err := o.merchantOrderRepository.GetByStatus(ctx, merchantID, statusReceive)
	if err != nil {
		return nil, err
	}

	if len(merchantOrder) == 0 {
		return &[]entities.MerchantOrder{}, nil
	}

	return &merchantOrder, nil
}
