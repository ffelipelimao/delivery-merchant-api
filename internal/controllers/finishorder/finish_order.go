package finishorder

import (
	"context"
	"net/http"

	"github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/dtos"
	interfaces "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/usecases"
	"github.com/gin-gonic/gin"
)

type FinishOrderHandler struct {
	FinishOrderUseCase interfaces.FinishOrderUseCase
}

func NewFinishOrderHandler(FinishOrderUseCase interfaces.FinishOrderUseCase) *FinishOrderHandler {
	return &FinishOrderHandler{
		FinishOrderUseCase: FinishOrderUseCase,
	}
}

// TODO swagger
func (h *FinishOrderHandler) Handle(c *gin.Context) {
	merchantID := c.Param("merchant_id")
	orderID := c.Param("order_id")

	var finishOrder dtos.FinishOrderDTO

	if err := c.BindJSON(&finishOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	finishOrder.MerchantID = merchantID
	finishOrder.OrderID = orderID

	ctx := context.Background()

	if err := h.FinishOrderUseCase.Handle(ctx, &finishOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
}
