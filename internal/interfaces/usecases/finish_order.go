package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
)

//go:generate mockgen -destination=./../../mocks/usecases/finish_order_mock.go -source=./finish_order.go -package=mocks_usecases
type FinishOrderUseCase interface {
	Handle(ctx context.Context, input *dtos.FinishOrderDTO) error
}
