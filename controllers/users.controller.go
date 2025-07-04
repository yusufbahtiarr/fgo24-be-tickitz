package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get User Profile
// @Description  Get detailed profile information of the currently logged-in user based on the JWT token.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response{results=models.UserProfile}
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /user/profile [get]
func GetUserProfileHandler(ctx *gin.Context) {
	userIdx, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized: user ID not found in context",
		})
		return
	}

	userIdStr, ok := userIdx.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Invalid user ID format in token",
		})
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to parse user ID",
		})
		return
	}

	// role, _ := ctx.Get("role")

	// roleStr, ok := role.(string)
	// if !ok || roleStr != "user" {
	// 	ctx.JSON(http.StatusForbidden, utils.Response{
	// 		Success: false,
	// 		Message: "Forbidden: Access is allowed for 'user' role only",
	// 	})
	// 	return
	// }

	user, err := models.FindUserProfile(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success View User Profile",
		Results: user,
	})

}

// @Summary      Edit User Profile
// @Description  Edit user profile based on the JWT token provided.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        data  body      dto.UpdateProfileRequest  true  "Updated profile data"
// @Success      200   {object}  utils.Response
// @Failure      400   {object}  utils.Response
// @Failure      401   {object}  utils.Response
// @Failure      500   {object}  utils.Response
// @Security     BearerAuth
// @Router       /user/profile [patch]
func EditProfileHandler(ctx *gin.Context) {
	userIdx, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized: user ID not found in context",
		})
		return
	}

	userIdStr, ok := userIdx.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Invalid user ID format in token",
		})
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to parse user ID",
		})
		return
	}

	newData := dto.UpdateProfileRequest{}
	err = ctx.ShouldBindJSON(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input.",
			Errors:  err.Error(),
		})
		return
	}
	err = models.UpdateProfile(userId, newData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Update User Profile",
	})
}

// @Summary      Get User Transaction History
// @Description  Retrieve the transaction history of the currently logged-in user based on the JWT token.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response{results=[]models.Transaction}  "Successfully retrieved transaction history"
// @Failure      401  {object}  utils.Response  "Unauthorized: user ID not found or invalid token"
// @Failure      500  {object}  utils.Response  "Internal server error"
// @Security     BearerAuth
// @Router       /user/transaction-history [get]
func GetTransactionHistoryHandler(ctx *gin.Context) {
	userIdx, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized: user ID not found",
		})
		return
	}

	userIdStr, ok := userIdx.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Invalid user ID format in token",
		})
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to parse user ID",
		})
		return
	}

	history, err := models.GetTransactionHistory(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success show history transaction",
		Results: history,
	})
}
