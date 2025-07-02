package routers

import (
	"fgo24-be-tickitz/controllers"

	"github.com/gin-gonic/gin"
)

func movieRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GetListMovies)
	r.GET("/:id", controllers.GetMovieByID)
	r.GET("/upcoming", controllers.GetUpcomingMovies)
}
