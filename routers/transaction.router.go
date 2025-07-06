package routers

import (
	"fgo24-be-tickitz/controllers"
	"fgo24-be-tickitz/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("/:id/ticket", controllers.GetTicketResultHandler)
	r.GET("/booked-seats", controllers.GetBookedSeatsInfoHandler)
	r.POST("", controllers.CreateTransactionHandler)
}
