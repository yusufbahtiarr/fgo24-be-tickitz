package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
