package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(ctx *gin.Context) {
	reqUser := dto.RegisterUserRequest{}
	if err := ctx.ShouldBindJSON(&reqUser); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
			Errors:  err.Error(),
		})
		return
	}

	if reqUser.Password != reqUser.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Password and Confirm Password do not match",
		})
		return
	}

	exist, err := models.CheckUserExistsByEmail(reqUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to verify email existence",
			Errors:  err.Error(),
		})
		return
	}

	if exist {
		ctx.JSON(http.StatusConflict, utils.Response{
			Success: false,
			Message: fmt.Sprintf("Email %s is already registered", reqUser.Email),
		})
		return
	}

	err = models.RegisterUser(reqUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed register user",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Register User",
		Results: reqUser.Email,
	})

}

func LoginUserHandler(ctx *gin.Context) {
	godotenv.Load()
	secretKey := os.Getenv("APP_SECRET")
	loginUser := dto.LoginUserRequest{}

	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	user, err := models.FindUserByEmail(loginUser.Email)
	if err != nil {
		fmt.Println(user)
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User with specified email not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{ // Bisa juga http.StatusUnauthorized
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	})

	token, _ := generatedToken.SignedString([]byte(secretKey))

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Login.",
		Results: map[string]string{
			"token": token,
		},
	})
}
