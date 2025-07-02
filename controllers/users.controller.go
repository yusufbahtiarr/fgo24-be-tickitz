package controllers

import (
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"

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
	userIdx, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized: No user ID found in token",
		})
		return
	}

	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "user" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'user' role only",
		})
		return
	}

	userIdFloat, ok := userIdx.(float64)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to convert user ID",
		})
		return
	}
	userId := int(userIdFloat)

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
