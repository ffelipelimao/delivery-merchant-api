#!/bin/bash

echo "Criando a fila SQS..."
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name production_start_queue --region us-east-1
