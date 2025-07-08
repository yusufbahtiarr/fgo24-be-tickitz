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
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Register User
// @Description  Register a new user with email and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      dto.RegisterUserRequest  true  "User Registration Input"
// @Success      200   {object}  utils.Response
// @Failure      400   {object}  utils.Response
// @Failure      409   {object}  utils.Response
// @Failure      500   {object}  utils.Response
// @Router       /auth/register [post]
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

// @Summary      Login User
// @Description  Authenticate user with email and password, and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      dto.LoginUserRequest  true  "User login credentials"
// @Success      200   {object}  utils.Response
// @Failure      400   {object}  utils.Response
// @Failure      404   {object}  utils.Response
// @Failure      401   {object}  utils.Response
// @Router       /auth/login [post]
func LoginUserHandler(ctx *gin.Context) {
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
			Message: "Email is not registered",
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

	token := GeneratedToken(user)

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Login.",
		Results: map[string]string{
			"token": "Bearer " + token,
		},
	})
}

// @Summary      Forgot Password
// @Description  Request password reset by verifying email and returning reset token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email  body      dto.ForgotPasswordRequest  true  "User email for password reset"
// @Success      200    {object}  utils.Response
// @Failure      400    {object}  utils.Response
// @Failure      404    {object}  utils.Response
// @Router       /auth/forgot-password [post]
func ForgotPasswordHandler(ctx *gin.Context) {
	forgotPassword := dto.ForgotPasswordRequest{}

	if err := ctx.ShouldBindJSON(&forgotPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	checkEmail, err := models.FindUserByEmail(forgotPassword.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Email not registered",
		})
		return
	}

	token := GeneratedResetPasswordToken(checkEmail)

	err = utils.SendEmail(forgotPassword.Email, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to send email",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success request forgot password. Check your email.",
	})
}

// @Summary      Reset Password
// @Description  Reset user password using a valid reset token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body   body      dto.ResetPasswordRequest   true  "New password data"
// @Success      200    {object}  utils.Response
// @Failure      400    {object}  utils.Response
// @Failure      401    {object}  utils.Response
// @Failure      404    {object}  utils.Response
// @Router       /auth/reset-password [post]
func ResetPasswordHandler(ctx *gin.Context) {
	resPassword := dto.ResetPasswordRequest{}

	if err := ctx.ShouldBindJSON(&resPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Request",
		})
		return
	}

	userID, err := models.VerifyResetPasswordToken(resPassword.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	_, err = models.FindUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response{
			Success: false,
			Message: "User does not exist",
		})
		return
	}

	err = models.ResetPassword(userID, resPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to reset password.",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Password reset successful.",
	})
}

func GeneratedToken(user models.User) string {
	secretKey := os.Getenv("APP_SECRET")
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(1 * time.Hour).Unix(),
	})

	token, _ := generatedToken.SignedString([]byte(secretKey))

	return token
}

func GeneratedResetPasswordToken(user models.User) string {
	secretKey := os.Getenv("APP_SECRET")
	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  user.ID,
		"purpose": "reset_password",
		"exp":     time.Now().Add(10 * time.Minute).Unix(),
	})

	token, _ := generatedToken.SignedString([]byte(secretKey))

	return token
}
