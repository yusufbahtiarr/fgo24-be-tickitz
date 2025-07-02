package routers

import (
	"fgo24-be-tickitz/controllers"
	"fgo24-be-tickitz/middlewares"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("/profile", middlewares.VerifyToken(), controllers.GetUserProfileHandler)
	// r.POST("/profile", middlewares.VerifyToken(), controller.EditProfile)
}
