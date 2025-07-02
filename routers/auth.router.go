package routers

import (
	"fgo24-be-tickitz/controllers"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.POST("/register", controllers.RegisterUserHandler)
	r.POST("/login", controllers.LoginUserHandler)
	r.POST("/forgot-password", controllers.ForgotPasswordHandler)
}
