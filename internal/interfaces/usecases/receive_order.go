package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
)

//go:generate mockgen -destination=./../../mocks/usecases/receive_order_mock.go -source=./receive_order.go -package=mocks_usecases
type ReceiveOrderUseCase interface {
	Handle(ctx context.Context, input *dtos.ReceiveOrderDTO) error
}
