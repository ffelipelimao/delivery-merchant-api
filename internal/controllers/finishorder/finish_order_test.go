package finishorder_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/controllers/finishorder"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	mocks_usecases "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/usecases"
)

func Test_IT_FinishOrder_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           string
		merchantID     string
		orderID        string
		mock           func(m *mocks_usecases.MockFinishOrderUseCase)
		expectedStatus int
	}{
		{
			name:       "successful finish order",
			body:       `{"status":"Finalizado"}`,
			merchantID: "123",
			orderID:    "456",
			mock: func(m *mocks_usecases.MockFinishOrderUseCase) {
				m.EXPECT().Handle(gomock.Any(), &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"}).Return(nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "should return 4xx invalid body",
			body:           `{`,
			merchantID:     "123",
			orderID:        "456",
			mock:           func(m *mocks_usecases.MockFinishOrderUseCase) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:       "should return 5xx invalid error to handle",
			body:       `{"status":"Finalizado"}`,
			merchantID: "123",
			orderID:    "456",
			mock: func(m *mocks_usecases.MockFinishOrderUseCase) {
				m.EXPECT().Handle(gomock.Any(), &dtos.FinishOrderDTO{MerchantID: "123", OrderID: "456", Status: "Finalizado"}).Return(assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			defer ctrl.Finish()
			finishOrderUseCaseMock := mocks_usecases.NewMockFinishOrderUseCase(ctrl)

			tc.mock(finishOrderUseCaseMock)

			router := gin.New()

			finishOrderHandler := finishorder.NewFinishOrderHandler(finishOrderUseCaseMock)

			router.POST("/merchant/:merchant_id/order/:order_id/finish", finishOrderHandler.Handle)

			reqBody := bytes.NewBufferString(tc.body)
			req, _ := http.NewRequest(http.MethodPost, "/merchant/"+tc.merchantID+"/order/"+tc.orderID+"/finish", reqBody)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}
