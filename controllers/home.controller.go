package controllers

import (
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

// @Summary      Get Upcoming Movies
// @Description  Retrieve a list of upcoming movies
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.UpcomingMovie
// @Failure      500  {object}  utils.Response
// @Router       / [get]
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

// @Summary      List Movies
// @Description  Get list of movies with optional search by title
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Param        search  query     string  false  "Search by movie title"
// @Success      200     {array}  models.Movies
// @Failure      500     {object}  utils.Response
// @Router       /movies [get]
func GetListMovies(ctx *gin.Context) {
	searchTitle := ctx.Query("search")
	movies, err := models.FindListMovies(searchTitle)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, utils.Response{
				Success: false,
				Message: "no movies matching the search criteria",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to search movies by title",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Movies",
		Results: movies,
	})
}

// @Summary      Get Movie By ID
// @Description  Get movie detail by movie ID
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  models.Movie
// @Failure      404  {object}  utils.Response
// @Router       /buy-ticket/{id} [get]
func GetMovieByID(ctx *gin.Context) {
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
		})
		return
	}

	movie, err := models.FindMovieByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to show movies by id",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Movie By ID",
		Results: movie,
	})
}
