package routers

import (
	"fgo24-be-tickitz/controllers"
	"fgo24-be-tickitz/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineGroup(r *gin.Engine) {
	authRouter(r.Group("/auth"))
	userRouter(r.Group("/users"))
	r.GET("", controllers.GetUpcomingMovies)
	r.GET("/movies", controllers.GetListMovies)
	r.GET("/buy-ticket/:id", controllers.GetMovieByID)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
