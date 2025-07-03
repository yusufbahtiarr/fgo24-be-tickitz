package routers

import (
	"fgo24-be-tickitz/controllers"
	"fgo24-be-tickitz/middlewares"

	"github.com/gin-gonic/gin"
)

func adminRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("/movies", controllers.CreateMovieHandler)
	r.GET("/movies", controllers.GetAllMoviesCreatedHandler)

}
