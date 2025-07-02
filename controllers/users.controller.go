package controllers

import (
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get User Profile
// @Description  Get detailed information about the currently logged-in user.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response{results=models.UserProfile}
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /users/profile [get]
func GetUserProfileHandler(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	role, _ := ctx.Get("role")
	userIds := userIdx.(string)

	if role != "user" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'user' role only",
		})
		ctx.Abort()
		return
	}

	userId, err := strconv.Atoi(userIds)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Invalid user ID format",
		})
		ctx.Abort()
		return
	}

	// fmt.Printf("User yang sedang login adalah user dengan id %d dengan role %s\n", userId, role)

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
