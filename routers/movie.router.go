package routers

import (
	"fgo24-be-tickitz/controllers"

	"github.com/gin-gonic/gin"
)

func movieRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GetListMoviesHandler)
	r.GET("/:id", controllers.GetMovieByIDHandler)
	r.GET("/upcoming", controllers.GetUpcomingMoviesHandler)
	r.GET("/now-showing", controllers.GetNowShowingMoviesHandler)
}
