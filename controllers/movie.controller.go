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
// @Router       /movies/upcoming [get]
func GetUpcomingMoviesHandler(ctx *gin.Context) {
	movies, err := models.GetUpcomingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Upcoming Movies",
		Results: movies,
	})
}

// @Summary      Get Now Showing Movies
// @Description  Retrieve a list of now showing movies
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.NowShowingMovie
// @Failure      500  {object}  utils.Response
// @Router       /movies/now-showing [get]
func GetNowShowingMoviesHandler(ctx *gin.Context) {
	movies, err := models.GetNowShowingMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Now Showing Movies",
		Results: movies,
	})
}

// @Summary      Get list of movies
// @Description  Retrieve a list of movies with optional filters such as title search, sorting, and pagination.
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Param        search query    string  false  "Search movies by title"
// @Param        genre query    string  false  "Filter movies by genre"
// @Param        sort query    string  false  "Sorting by field (e.g., popular, latest, title_asc, title_desc)"
// @Param        page   query    int     false  "Page"
// @Success      200    {object} utils.Response{results=[]models.Movie,page_info=utils.PageInfo}
// @Failure      400    {object} utils.Response
// @Failure      404    {object} utils.Response
// @Failure      500    {object} utils.Response
// @Router       /movies [get]
func GetListMoviesHandler(ctx *gin.Context) {
	searchTitle := ctx.Query("search")
	genre := ctx.Query("genre")
	sort := ctx.Query("sort")
	pageX := ctx.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pageX)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid page number. Page must be a positive integer.",
		})
		return
	}

	limit := 10
	offset := (page - 1) * limit

	movies, totalMovies, err := models.GetListMovies(searchTitle, genre, sort, limit, offset)
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
		return
	}

	totalPages := (totalMovies + int(limit) - 1) / int(limit)

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Movies",
		Results: movies,
		PageInfo: &utils.PageInfo{
			Total:      totalMovies,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
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
// @Failure      500  {object}  utils.Response
// @Router       /movies/{id} [get]
func GetMovieByIDHandler(ctx *gin.Context) {
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
		})
		return
	}

	movie, err := models.GetMovieByID(id)
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
