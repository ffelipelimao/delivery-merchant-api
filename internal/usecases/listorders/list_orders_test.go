package listorders_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	mocks_repository "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/repository"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/listorders"
)

func TestListOrdersUseCase_Handle(t *testing.T) {
	tests := []struct {
		name           string
		mock           func(m *mocks_repository.MockMerchantOrderRepository)
		merchantID     string
		expectedResult *[]entities.MerchantOrder
		expectedErr    error
	}{
		{
			name: "successful retrieval of orders",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				orders := []entities.MerchantOrder{{MerchantID: "1", Status: "Pago"}, {MerchantID: "2", Status: "Pago"}}
				m.EXPECT().GetByStatus(gomock.Any(), "123", "Pago").Return(orders, nil)
			},
			merchantID:     "123",
			expectedResult: &[]entities.MerchantOrder{{MerchantID: "1", Status: "Pago"}, {MerchantID: "2", Status: "Pago"}},
			expectedErr:    nil,
		},
		{
			name: "no orders found",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				var orders []entities.MerchantOrder
				m.EXPECT().GetByStatus(gomock.Any(), "123", "Pago").Return(orders, nil)
			},
			merchantID:     "123",
			expectedResult: &[]entities.MerchantOrder{},
			expectedErr:    nil,
		},
		{
			name: "database error during retrieval",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				m.EXPECT().GetByStatus(gomock.Any(), "123", "Pago").Return(nil, errors.New("database error"))
			},
			merchantID:     "123",
			expectedResult: nil,
			expectedErr:    errors.New("database error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks_repository.NewMockMerchantOrderRepository(ctrl)
			useCase := listorders.NewListOrdersUseCase(mockRepo)

			tc.mock(mockRepo)

			result, err := useCase.Handle(context.Background(), tc.merchantID)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
