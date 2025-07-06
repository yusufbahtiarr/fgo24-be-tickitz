package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fmt"
	"time"
)

type TicketResult struct {
	MovieTitle   string    `json:"movie_title"`
	MovieDate    time.Time `json:"movie_date"`
	ShowTime     time.Time `json:"show_time"`
	SeatCount    int       `json:"seat_count"`
	Seats        string    `json:"seats"`
	TotalPayment int       `json:"total_payment"`
}

type Seat string

func GetTicketResult(transactionId, userId int) (TicketResult, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return TicketResult{}, fmt.Errorf("failed to connect database")
	}
	defer conn.Close()

	query := `
	SELECT 
		m.title,
		t.movie_date,
		tm.time,
		COUNT(td.seat) AS seat_count,
		STRING_AGG(td.seat, ', ') AS seats,
		t.total_payment
	FROM transactions t
	JOIN movies m ON m.id = t.id_movie
	JOIN times tm ON tm.id = t.id_time
	JOIN transaction_details td ON td.id_transaction = t.id
	WHERE t.id_user = $1 AND t.id = $2
	GROUP BY m.title, t.movie_date, tm.time, t.total_payment;
	`
	var ticket TicketResult
	err = conn.QueryRow(context.Background(), query, userId, transactionId).Scan(
		&ticket.MovieTitle,
		&ticket.MovieDate,
		&ticket.ShowTime,
		&ticket.SeatCount,
		&ticket.Seats,
		&ticket.TotalPayment,
	)
	if err != nil {
		return TicketResult{}, fmt.Errorf("ticket not found (%w)", err)
	}

	return ticket, nil
}

func GetBookedSeatsInfo(movieID int, date string, timeID, locationID, cinemaID int) ([]string, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := `
		SELECT td.seat
		FROM transaction_details td
		JOIN transactions t ON td.id_transaction = t.id
		WHERE
			t.id_movie = $1 AND
			t.movie_date = $2 AND
			t.id_time = $3 AND
			t.id_location = $4 AND
			t.id_cinema = $5;
	`

	rows, err := conn.Query(context.Background(), query, movieID, date, timeID, locationID, cinemaID)
	if err != nil {
		return nil, fmt.Errorf("failed to query booked seats: %w", err)
	}
	defer rows.Close()

	var seats []string
	for rows.Next() {
		var seat string
		if err := rows.Scan(&seat); err != nil {
			return nil, fmt.Errorf("failed to scan seat: %w", err)
		}
		seats = append(seats, seat)
	}

	if len(seats) == 0 {
		return nil, fmt.Errorf("no booked seats found")
	}

	return seats, nil
}
