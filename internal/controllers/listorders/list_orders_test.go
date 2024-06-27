package listorders_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/controllers/listorders"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/entities"
	mocks "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/usecases"
)

func Test_IT_ListOrdersHandler_Handle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		merchantID     string
		mock           func(m *mocks.MockListOrdersUseCase)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:       "successful retrieval of orders",
			merchantID: "123",
			mock: func(m *mocks.MockListOrdersUseCase) {
				mockList := &[]entities.MerchantOrder{}
				m.EXPECT().Handle(context.Background(), "123").Return(mockList, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   "[]",
		},
		{
			name:       "error on order retrieval",
			merchantID: "123",
			mock: func(m *mocks.MockListOrdersUseCase) {
				m.EXPECT().Handle(context.Background(), "123").Return(nil, assert.AnError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"message":"assert.AnError general error for testing"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUseCase := mocks.NewMockListOrdersUseCase(ctrl)
			tc.mock(mockUseCase)

			handler := listorders.NewListOrdersHandler(mockUseCase)

			router := gin.New()
			router.GET("/merchant/:merchant_id/orders", handler.Handle)

			req, _ := http.NewRequest(http.MethodGet, "/merchant/"+tc.merchantID+"/orders", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedBody, rec.Body.String())
		})
	}
}
