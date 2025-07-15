package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddlewate() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{"http://146.190.102.54:9702", "http://localhost:8888", "http://localhost:5173"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	})
}
