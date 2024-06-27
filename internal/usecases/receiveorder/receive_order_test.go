package receiveorder_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/receiveorder"

	mocks_repository "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/repository"
)

func TestReceiveOrderUseCase_Handle(t *testing.T) {
	tests := []struct {
		name        string
		mock        func(m *mocks_repository.MockMerchantOrderRepository)
		input       *dtos.ReceiveOrderDTO
		expectedErr error
	}{
		{
			name: "successfully receives an order",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				merchantOrder := entities.NewMerchantOrder("123", "received", "456")
				m.EXPECT().Insert(gomock.Any(), merchantOrder).Return(nil)
			},
			input:       &dtos.ReceiveOrderDTO{MerchantID: "123", OrderID: "456", Status: "received"},
			expectedErr: nil,
		},
		{
			name: "fails to insert order due to database error",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				merchantOrder := entities.NewMerchantOrder("123", "received", "456")
				m.EXPECT().Insert(gomock.Any(), merchantOrder).Return(errors.New("database error"))
			},
			input:       &dtos.ReceiveOrderDTO{MerchantID: "123", OrderID: "456", Status: "received"},
			expectedErr: errors.New("database error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks_repository.NewMockMerchantOrderRepository(ctrl)
			useCase := receiveorder.NewReceiveOrderUseCase(mockRepo)

			tc.mock(mockRepo)

			err := useCase.Handle(context.Background(), tc.input)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
