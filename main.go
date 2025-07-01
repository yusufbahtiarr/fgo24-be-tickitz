package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	godotenv.Load()
	r.Run(":" + os.Getenv("APP_PORT"))
}
