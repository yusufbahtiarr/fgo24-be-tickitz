package middlewares

import (
	"fgo24-be-tickitz/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		secretKey := os.Getenv("APP_SECRET")
		token := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")

		if len(token) < 2 {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		rawTokenString := strings.TrimSpace(token[1])
		rawToken, err := jwt.Parse(rawTokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !rawToken.Valid {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid or expired token",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId := rawToken.Claims.(jwt.MapClaims)["userId"]
		role := rawToken.Claims.(jwt.MapClaims)["role"]
		ctx.Set("userId", userId)
		ctx.Set("role", role)
		ctx.Next()
	}
}
