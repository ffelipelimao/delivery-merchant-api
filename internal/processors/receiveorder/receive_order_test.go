package receiveorder_test

import (
	"context"
	"testing"
	"time"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/external/queue"
	mocks_queue "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/queue"
	mocks_usecases "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/mocks/usecases"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/processors/receiveorder"

	"go.uber.org/mock/gomock"
)

func TestProcessor_Handle(t *testing.T) {
	tests := []struct {
		name string
		mock func(mq *mocks_queue.MockQueue, mu *mocks_usecases.MockReceiveOrderUseCase)
	}{
		{
			name: "successful message processing",
			mock: func(mq *mocks_queue.MockQueue, mu *mocks_usecases.MockReceiveOrderUseCase) {
				messages := []queue.Content{
					{Message: `{"merchantID":"123","orderID":"456","status":"received"}`, ReceiptHandle: "handle1"},
				}
				mq.EXPECT().PollingMessages().Return(messages, nil).AnyTimes()
				mq.EXPECT().DeleteMessageFromQueue("handle1").Return(nil).AnyTimes()

				mu.EXPECT().Handle(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockQueue := mocks_queue.NewMockQueue(ctrl)
			mockUseCase := mocks_usecases.NewMockReceiveOrderUseCase(ctrl)

			processor := receiveorder.NewProcessor(mockQueue, mockUseCase)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			tt.mock(mockQueue, mockUseCase)
			go processor.Handle(ctx)

			// Wait the goroutine process, is to avoid race conditions
			time.Sleep(1 * time.Second)
		})
	}
}
