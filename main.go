package main

import (
	"fgo24-be-tickitz/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @Title 			Tickitz - Booking Ticket App
// @Version 		1.0
// @Description	API Documentation Booking Ticket Application
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	r := gin.Default()
	godotenv.Load()
	routers.CombineGroup(r)
	r.Run(":" + os.Getenv("APP_PORT"))
}
