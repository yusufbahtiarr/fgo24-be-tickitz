package controllers

import (
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUpcomingMovies(ctx *gin.Context) {
	movies, err := models.FindUpcomingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Upcoming Movies",
		Results: movies,
	})
}

func GetListMovies(ctx *gin.Context) {
	movies, err := models.FindListMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Movies",
		Results: movies,
	})
}
