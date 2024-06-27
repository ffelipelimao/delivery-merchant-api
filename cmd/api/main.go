package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/config"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/controllers"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/controllers/finishorder"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/controllers/listorders"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/external/db"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/external/queue"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/processors/receiveorder"
	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/repository"
	finishorderusecase "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/finishorder"
	listordersusecase "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/listorders"
	receiveorderderusecase "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/usecases/receiveorder"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := config.NewConfig()
	err := configs.Load()
	if err != nil {
		log.Panic("error to init config", err)
	}

	db, err := db.NewDatabase(configs)
	if err != nil {
		log.Panic("error to init db", err)
	}

	queue := queue.NewQueue(configs)

	merchantOrderRepository := repository.NewMerchantRepository(db)

	receiveOrderUseCase := receiveorderderusecase.NewReceiveOrderUseCase(merchantOrderRepository)
	receveiOrderProcessor := receiveorder.NewProcessor(queue, receiveOrderUseCase)

	go receveiOrderProcessor.Handle(context.Background())

	finishOrderUseCase := finishorderusecase.NewFinishOrderUseCase(merchantOrderRepository)
	finishOrderHandler := finishorder.NewFinishOrderHandler(finishOrderUseCase)

	listOrdersUseCase := listordersusecase.NewListOrdersUseCase(merchantOrderRepository)
	listOrdersHandler := listorders.NewListOrdersHandler(listOrdersUseCase)

	app := gin.Default()
	router := controllers.Router{
		FinishOrderHandler: finishOrderHandler.Handle,
		ListOrdersHandler:  listOrdersHandler.Handle,
	}

	app.GET("/healthcheck", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Register(app)

	app.Run(":8080")
}
