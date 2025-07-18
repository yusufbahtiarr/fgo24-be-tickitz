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
	r.GET("/genres", controllers.GetGenresHandler)
	r.GET("/times", controllers.GetTimesHandler)
	r.GET("/locations", controllers.GetLocationsHandler)
	r.GET("/cinemas", controllers.GetCinemasHandler)
	r.GET("/payment-methods", controllers.GetPaymentMethodsHandler)
	r.GET("/casts", controllers.GetCastsHandler)
	r.GET("/directors", controllers.GetDirectorsHandler)
}
