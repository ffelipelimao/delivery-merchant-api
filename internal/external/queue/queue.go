package queue

import (
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Queue struct {
	svc      *sqs.SQS
	queueUrl string
}

type Content struct {
	Message       string
	ReceiptHandle string
}

func NewQueue(config *config.Config) *Queue {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint: aws.String("https://sqs.us-east-1.amazonaws.com"),
			Region:   aws.String("us-east-1"),
		},
	}))
	svc := sqs.New(sess)
	return &Queue{
		svc:      svc,
		queueUrl: config.QueueUrl,
	}
}

func (q *Queue) PollingMessages() ([]Content, error) {
	result, err := q.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        aws.String(q.queueUrl),
		WaitTimeSeconds: aws.Int64(20),
	})
	if err != nil {
		return nil, err
	}

	messagesContents := make([]Content, 0, len(result.Messages))

	for _, message := range result.Messages {
		content := Content{
			Message:       aws.StringValue(message.Body),
			ReceiptHandle: aws.StringValue(message.ReceiptHandle),
		}
		messagesContents = append(messagesContents, content)
	}
	return messagesContents, nil
}

func (q *Queue) DeleteMessageFromQueue(receiptHandle string) error {
	_, err := q.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(q.queueUrl),
		ReceiptHandle: aws.String(receiptHandle),
	})
	if err != nil {
		return err
	}
	return nil
}
