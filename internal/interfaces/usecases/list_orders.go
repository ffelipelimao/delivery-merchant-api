package interfaces

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/usecases/list_orders_mock.go -source=./list_orders.go -package=mocks_usecases
type ListOrdersUseCase interface {
	Handle(ctx context.Context, merchantID string) (*[]entities.MerchantOrder, error)
}
