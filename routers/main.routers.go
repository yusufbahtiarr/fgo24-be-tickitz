package routers

import (
	"fgo24-be-tickitz/controllers"

	"github.com/gin-gonic/gin"
)

func CombineGroup(r *gin.Engine) {
	r.GET("", controllers.GetUpcomingMovies)
	r.GET("/movies", controllers.GetListMovies)
}
