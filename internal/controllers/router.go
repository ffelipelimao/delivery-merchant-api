package controllers

import "github.com/gin-gonic/gin"

type Router struct {
	FinishOrderHandler gin.HandlerFunc
	ListOrdersHandler  gin.HandlerFunc
}

func (r *Router) Register(app *gin.Engine) {
	delivery := app.Group("/v1/delivery-merchant")
	{
		delivery.PATCH("/:merchant_id/orders/:order_id", r.FinishOrderHandler)
		delivery.GET("/:merchant_id/orders", r.ListOrdersHandler)
	}
}
