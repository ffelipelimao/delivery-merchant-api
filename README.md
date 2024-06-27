# delivery-merchant-api


### Subir aplicação

- Subir o docker-compose com o run.sh
- Criar a tabela dentro de docs/db
- Criar a fila dentro de docs/queue/create_queue.sh

*Não se esqueça de dar o chmod +x para executar os .sh*

### Testar local a aplicação

- Altere os valores na string $MENSAGEM
- Criar uma mensagem  de docs/queue/publish_queue.sh 

*Não se esqueça de dar o chmod +x para executar os .sh*

#### Endpoints

- GET `/v1/delivery-merchant/:id/orders` para lista todos os pedidos "Pago"
- PATCH `/v1/delivery-merchant/:id/orders/:order_id` para finalizar um pedido "Finalizado"


#### Test

- go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html && open coverage.html



#### Deploy

- Criar o cluster na AWS
```eksctl create cluster --name=delivery-merchant-cluster --region=us-east-1 --node-type=t3.micro```
- Criar o RDS
- Criar a fila SQS
- Após criar o cluster, dar permissão no usuario IAM para ler a mensagem
- Atualizar a VPC, sub-redes e o internet gateway pra aceitar qualquer conexão
- Apos fazer um deploy, matar os pods antigos para que os novos vão subir