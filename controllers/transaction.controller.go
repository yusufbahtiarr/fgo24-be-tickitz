package controllers

import (
	"fgo24-be-tickitz/dto"
	"fgo24-be-tickitz/models"
	"fgo24-be-tickitz/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get Ticket Result by Transaction ID
// @Description  Retrieve the ticket details for a specific transaction ID, accessible only by the authenticated user.
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        transaction_id  query     int  true  "Transaction ID"
// @Success      200  {object}   utils.Response{results=models.TicketResult}
// @Failure      400  {object}   utils.Response
// @Failure      401  {object}   utils.Response
// @Failure      404  {object}   utils.Response
// @Failure      500  {object}   utils.Response
// @Security     BearerAuth
// @Router       /transactions/{transaction_id}/ticket [get]
func GetTicketResultHandler(ctx *gin.Context) {
	transactionIdx := ctx.Query("transaction_id")
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

	transactionId, err := strconv.Atoi(transactionIdx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to parse transaction ID",
		})
		return
	}

	ticket, err := models.GetTicketResult(transactionId, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "failed to get ticket user",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "success show ticket result user",
		Results: ticket,
	})
}

// @Summary      Get Booked Seats Info
// @Description  Retrieve a list of already booked seats based on movie, date, time, location, and studio.
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        movie_id     query    int     true  "Movie ID"
// @Param        date         query    string  true  "Movie date in format YYYY-MM-DD"
// @Param        time_id      query    int     true  "Time ID"
// @Param        location_id  query    int     true  "Location ID"
// @Param        cinema_id    query    int     true  "Cinema ID"
// @Success      200  {object}  utils.Response{results=[]string}
// @Failure      400  {object}  utils.Response
// @Failure      401  {object}  utils.Response
// @Failure      500  {object}  utils.Response
// @Security     BearerAuth
// @Router       /transactions/booked-seats [get]
func GetBookedSeatsInfoHandler(ctx *gin.Context) {
	movieIdX := ctx.Query("movie_id")
	date := ctx.Query("date")
	timeIdX := ctx.Query("time_id")
	locationIdX := ctx.Query("location_id")
	cinemaIdX := ctx.Query("cinema_id")

	movieId, err := strconv.Atoi(movieIdX)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie_id",
		})
		return
	}

	timeId, err := strconv.Atoi(timeIdX)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid time_id",
		})
		return
	}

	locationId, err := strconv.Atoi(locationIdX)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid location_id",
		})
		return
	}

	cinemaId, err := strconv.Atoi(cinemaIdX)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid cinema_id",
		})
		return
	}

	_, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized: user ID not found",
		})
		return
	}

	seats, err := models.GetBookedSeatsInfo(movieId, date, timeId, locationId, cinemaId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to show booked seat info",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success get booked seat info",
		Results: seats,
	})
}

// @Summary      Create new transaction
// @Description  Create a new transaction with seat selection for movie booking
// @Tags         Transactions
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateTransactionRequest true "Transaction payload"
// @Success      201 {object} utils.Response
// @Failure      400 {object} utils.Response
// @Failure      401 {object} utils.Response
// @Failure      500 {object} utils.Response
// @Security     BearerAuth
// @Router       /transactions [post]
func CreateTransactionHandler(ctx *gin.Context) {
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

	transaction := dto.CreateTransactionRequest{}

	if err = ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request transaction",
		})
		return
	}

	err = models.CreateTransaction(userId, transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed Create Transaction",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Success Create Transaction",
		Results: transaction})
}
