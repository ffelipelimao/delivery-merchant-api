package queue

import "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/external/queue"

//go:generate mockgen -destination=./../../mocks/queue/queue_mock.go -source=./queue.go -package=mocks_queue
type Queue interface {
	PollingMessages() ([]queue.Content, error)
	DeleteMessageFromQueue(receiptHandle string) error
}
