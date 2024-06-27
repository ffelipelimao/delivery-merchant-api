package listorders

import (
	"context"
	"net/http"

	interfaces "github.com/Pos-Tech-Challenge-48/delivery-merchant-api/internal/interfaces/usecases"
	"github.com/gin-gonic/gin"
)

type ListOrdersHandler struct {
	ListOrdersUseCase interfaces.ListOrdersUseCase
}

func NewListOrdersHandler(ListOrdersUseCase interfaces.ListOrdersUseCase) *ListOrdersHandler {
	return &ListOrdersHandler{
		ListOrdersUseCase: ListOrdersUseCase,
	}
}

// TODO swagger
func (h *ListOrdersHandler) Handle(c *gin.Context) {
	merchantID := c.Param("merchant_id")

	ctx := context.Background()

	orders, err := h.ListOrdersUseCase.Handle(ctx, merchantID)
	if err != nil {
		// TODO error 404
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
