package finishorderusecase

import (
	"context"
	"errors"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/repository"
)

type FinishOrderUseCase struct {
	merchantOrderRepository repository.MerchantOrderRepository
}

func NewFinishOrderUseCase(merchantOrderRepository repository.MerchantOrderRepository) *FinishOrderUseCase {
	return &FinishOrderUseCase{
		merchantOrderRepository: merchantOrderRepository,
	}
}

func (o *FinishOrderUseCase) Handle(ctx context.Context, input *dtos.FinishOrderDTO) error {
	merchantOrder, err := o.merchantOrderRepository.Get(ctx, input.MerchantID, input.OrderID)
	if err != nil {
		return err
	}

	if merchantOrder == nil {
		return errors.New("not found")
	}

	err = merchantOrder.ValidateStatus(input.Status)
	if err != nil {
		return err
	}

	merchantOrder.Status = input.Status

	err = o.merchantOrderRepository.Update(ctx, merchantOrder)
	if err != nil {
		return err
	}
	return nil
}
