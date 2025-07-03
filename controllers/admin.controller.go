package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary     Create new movie
// @Description Create movie with genre, cast, and director
// @Tags        Admins
// @Accept      json
// @Produce     json
// @Param       request body dto.CreateMovieRequest true "Movie payload"
// @Success     201 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     500 {object} utils.Response
// @Security     BearerAuth
// @Router      /admins/movies [post]
func CreateMovieHandler(ctx *gin.Context) {
	movie := dto.CreateMovieRequest{}
	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request movie",
		})
		return
	}

	err := models.CreateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed Create Movie",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success Create Movie",
		Results: map[string]string{
			"title": movie.Title,
		}})
}

// @Summary     Get Movie Created
// @Description Get list all movies created through admin
// @Tags        Admins
// @Accept      json
// @Produce     json
// @Success     201 {object} utils.Response
// @Failure     400 {object} utils.Response
// @Failure     500 {object} utils.Response
// @Security     BearerAuth
// @Router      /admins/movies [get]
func GetAllMoviesCreatedHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}

	movies, err := models.GetAllMoviesCreated()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed list movies created",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: false,
		Message: "Success show list movies created",
		Results: movies,
	})

}
