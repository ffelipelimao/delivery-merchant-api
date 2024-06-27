package receiveorder

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	qinterfaces "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/queue"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/usecases"
)

type Processor struct {
	queue               qinterfaces.Queue
	receiveOrderUseCase interfaces.ReceiveOrderUseCase
}

func NewProcessor(queue qinterfaces.Queue, receiveOrderUseCase interfaces.ReceiveOrderUseCase) *Processor {
	return &Processor{
		queue:               queue,
		receiveOrderUseCase: receiveOrderUseCase,
	}
}

func (p *Processor) Handle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			messagesContent, err := p.queue.PollingMessages()
			if err != nil {
				fmt.Println("error to pooling message:", err.Error())
				continue
			}
			if messagesContent == nil {
				continue
			}
			for _, messageContent := range messagesContent {
				var receiveOrder dtos.ReceiveOrderDTO
				err := json.Unmarshal([]byte(messageContent.Message), &receiveOrder)
				if err != nil {
					fmt.Println("invalid message format:", err.Error())
					continue
				}

				ctx := context.Background()

				err = p.receiveOrderUseCase.Handle(ctx, &receiveOrder)
				if err != nil {
					fmt.Println("error to handle message:", err.Error())
					continue
				}

				err = p.queue.DeleteMessageFromQueue(messageContent.ReceiptHandle)
				if err != nil {
					fmt.Println("error to delete message:", err.Error())
				}
			}
		}
	}
}
