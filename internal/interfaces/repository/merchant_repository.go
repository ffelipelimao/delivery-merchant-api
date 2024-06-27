package repository

import (
	"context"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
)

//go:generate mockgen -destination=./../../mocks/repository/merchant_repository_mock.go -source=./merchant_repository.go -package=mocks_repository
type MerchantOrderRepository interface {
	Update(ctx context.Context, merchantOrder *entities.MerchantOrder) error
	Get(ctx context.Context, merchantID string, orderID string) (*entities.MerchantOrder, error)
	GetByStatus(ctx context.Context, merchantID string, status string) ([]entities.MerchantOrder, error)
	Insert(ctx context.Context, merchantOrder *entities.MerchantOrder) error
}
