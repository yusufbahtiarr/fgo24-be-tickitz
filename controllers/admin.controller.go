package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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
// @Router      /admin/movies [post]
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

// @Summary      Get All Movies Created by Admin
// @Description  Retrieve the list of all movies created through the admin panel
// @Tags         Admins
// @Accept       json
// @Produce      json
// @Param        release_month   query    string     false  "Filter by Release Month (e.g., 2025-07)"
// @Param        page   query    int     false  "Page (e.g., 1)"
// @Param        limit  query    int     false  "Items per page (e.g., 5)"
// @Success      200  {object}  utils.Response{results=[]models.CreatedMovies{},page_info=utils.PageInfo}
// @Failure      400  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /admin/movies [get]
func GetAllMoviesCreatedHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	releaseMonth := ctx.Query("release_month")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}

	pageX := ctx.DefaultQuery("page", "1")
	limitX := ctx.DefaultQuery("limit", "5")

	if releaseMonth != "" {
		matched, err := regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])$`, releaseMonth)
		if err != nil || !matched {
			ctx.JSON(http.StatusBadRequest, utils.Response{
				Success: false,
				Message: "Invalid Release Month format. Use YYYY-MM (e.g., 2025-06).",
			})
			return
		}
	}

	page, err := strconv.Atoi(pageX)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid page number. Page must be a positive integer.",
		})
		return
	}

	limit, err := strconv.Atoi(limitX)
	if err != nil || limit < 1 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid limit number. Limit must be a positive integer.",
		})
		return
	}

	offset := (page - 1) * limit

	movies, totalMovies, err := models.GetAllMoviesCreated(releaseMonth, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed list movies created",
			Errors:  err.Error(),
		})
		return
	}

	totalPages := (totalMovies + int(limit) - 1) / int(limit)

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success show list movies created",
		Results: movies,
		PageInfo: &utils.PageInfo{
			Total:      totalMovies,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	})

}

// @Summary      Get Movie by ID (Admin Only)
// @Description  Retrieve a specific movie created by the admin using its ID
// @Tags         Admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  utils.Response{results=models.DetailCreatedMovie}
// @Failure      400  {object}  utils.Response
// @Failure      403  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /admin/movies/{id} [get]
func GetMovieCreatedByIDHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
		})
		return
	}

	movie, err := models.GetMovieCreatedByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to show movies created by id",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "List Movie Created By ID",
		Results: movie,
	})
}

// @Summary      Delete Movie by ID (Admin Only)
// @Description  Delete a movie created by an admin using its ID
// @Tags         Admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      403  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /admin/movies/{id} [delete]
func DeleteMovieHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
		})
		return
	}

	err = models.DeleteMovie(id)
	if err != nil {
		if strings.Contains(err.Error(), "movie not found") {
			ctx.JSON(http.StatusNotFound, utils.Response{
				Success: false,
				Message: fmt.Sprintf("Movie with id = %d not found", id),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, utils.Response{
				Success: false,
				Message: "Failed to delete movie",
				Errors:  err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: fmt.Sprintf("Success delete movie with id = %d", id),
	})

}

// @Summary      Update Movie (Admin Only)
// @Description  Update a movie created by an admin using its ID
// @Tags         Admins
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Movie ID"
// @Param        body  body      dto.UpdateMovieRequest true  "Update movie body"
// @Success      200  {object}  utils.Response
// @Failure      400  {object}  utils.Response
// @Failure      403  {object}  utils.Response
// @Failure      404  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /admin/movies/{id} [patch]
func UpdateMovieHandler(ctx *gin.Context) {
	role, _ := ctx.Get("role")

	roleStr, ok := role.(string)
	if !ok || roleStr != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Access is allowed for 'admin' role only",
		})
		return
	}
	idx := ctx.Param("id")
	id, err := strconv.Atoi(idx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
		})
		return
	}

	newData := dto.UpdateMovieRequest{}
	err = ctx.ShouldBindJSON(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid input.",
			Errors:  err.Error(),
		})
		return
	}

	err = models.UpdateMovie(id, newData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update movie",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success update movie",
		Results: "Update Movie " + newData.Title,
	})
}
