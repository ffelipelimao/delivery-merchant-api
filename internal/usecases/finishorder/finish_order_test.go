package finishorderusecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	mocks_repository "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/repository"
	finishorderusecase "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/finishorder"
)

func TestFinishOrderUseCase_Handle(t *testing.T) {
	tests := []struct {
		name        string
		mock        func(m *mocks_repository.MockMerchantOrderRepository)
		input       *dtos.FinishOrderDTO
		expectedErr error
	}{
		{
			name: "successfully finishes order",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				mockOrder := &entities.MerchantOrder{Status: "received"}

				m.EXPECT().Get(gomock.Any(), "123", "456").Return(mockOrder, nil)
				m.EXPECT().Update(gomock.Any(), mockOrder).Return(nil)
			},
			input:       &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"},
			expectedErr: nil,
		},
		{
			name: "should return order not found",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				m.EXPECT().Get(gomock.Any(), "123", "456").Return(nil, nil)
			},
			input:       &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"},
			expectedErr: errors.New("not found"),
		},
		{
			name: "should be invalid changed",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				mockOrder := &entities.MerchantOrder{Status: "Finalizado"}

				m.EXPECT().Get(gomock.Any(), "123", "456").Return(mockOrder, nil)
			},
			input:       &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"},
			expectedErr: errors.New("order is already finish"),
		},
		{
			name: "should return an error to get an order",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				m.EXPECT().Get(gomock.Any(), "123", "456").Return(nil, assert.AnError)
			},
			input:       &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"},
			expectedErr: assert.AnError,
		},
		{
			name: "successfully return an error to update an order",
			mock: func(m *mocks_repository.MockMerchantOrderRepository) {
				mockOrder := &entities.MerchantOrder{Status: "received"}

				m.EXPECT().Get(gomock.Any(), "123", "456").Return(mockOrder, nil)
				m.EXPECT().Update(gomock.Any(), mockOrder).Return(assert.AnError)
			},
			input:       &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"},
			expectedErr: assert.AnError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks_repository.NewMockMerchantOrderRepository(ctrl)
			useCase := finishorderusecase.NewFinishOrderUseCase(mockRepo)

			tc.mock(mockRepo)

			err := useCase.Handle(context.TODO(), tc.input)
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
