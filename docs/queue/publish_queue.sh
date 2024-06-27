#!/bin/bash

MESSAGE='{"merchant_id": "a7eaf5d8-ac11-4768-948c-2ede10a881ff", "order_id": 67893, "status": "received"}'
QUEUE_URL="http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/production_start_queue"

aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url "$QUEUE_URL" --message-body "$MESSAGE"
echo "Mensagem enviada para a fila SQS: $MESSAGE"
