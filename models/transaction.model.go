package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fgo24-be-tickitz/dto"
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

type TransactionResult struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	TotalPayment  string        `json:"total_payment"`
	MovieDate     FormattedDate `json:"movie_date"`
	Title         string        `json:"title"`
	CinemaName    string        `json:"cinema_name"`
	Time          FormattedTime `json:"time"`
	Location      string        `json:"location"`
	PaymentMethod string        `json:"payment_method"`
}

type FormattedDate time.Time
type FormattedTime time.Time

func (d FormattedDate) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(formatted), nil
}

func (t FormattedTime) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	formatted := fmt.Sprintf("\"%s\"", tt.Format("15:04:05"))
	return []byte(formatted), nil
}

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

	return seats, nil
}

func CreateTransaction(userId int, NewTransaction dto.CreateTransactionRequest) (TransactionResult, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return TransactionResult{}, err
	}
	defer conn.Close()

	ctx := context.Background()

	if len(NewTransaction.Seats) == 0 {
		return TransactionResult{}, fmt.Errorf("seats cannot be empty")
	}
	if NewTransaction.TotalPayment <= 0 {
		return TransactionResult{}, fmt.Errorf("total payment must be greater than 0")
	}

	var bookedSeats []string
	queryCheck := `
		SELECT td.seat
		FROM transaction_details td
		JOIN transactions t ON t.id = td.id_transaction
		WHERE 
			t.id_movie = $1 AND	
			t.movie_date = $2 AND	
			t.id_time = $3 AND	
			t.id_cinema = $4 AND	
			t.id_location = $5 AND	
			td.seat = ANY($6::text[])
	`

	rows, err := conn.Query(ctx, queryCheck,
		NewTransaction.MovieID,
		NewTransaction.MovieDate,
		NewTransaction.TimeID,
		NewTransaction.CinemaID,
		NewTransaction.LocationID,
		NewTransaction.Seats,
	)
	if err != nil {
		return TransactionResult{}, fmt.Errorf("failed to check existing booked seats: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var seat string
		if err = rows.Scan(&seat); err != nil {
			return TransactionResult{}, fmt.Errorf("failed to find booked seat: %w", err)
		}
		bookedSeats = append(bookedSeats, seat)
	}

	if len(bookedSeats) > 0 {
		return TransactionResult{}, fmt.Errorf("seats already booked: %v", bookedSeats)
	}

	trx, err := conn.Begin(ctx)
	if err != nil {
		return TransactionResult{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	succesed := false
	defer func() {
		if !succesed {
			trx.Rollback(ctx)
		}
	}()

	query := `
		INSERT INTO transactions 
		(name, email, phone, total_payment, movie_date, id_movie, id_cinema, id_time, id_location, id_payment_method, id_user)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;
	`
	var transactionID int
	err = trx.QueryRow(ctx, query,
		NewTransaction.Name,
		NewTransaction.Email,
		NewTransaction.Phone,
		NewTransaction.TotalPayment,
		NewTransaction.MovieDate,
		NewTransaction.MovieID,
		NewTransaction.CinemaID,
		NewTransaction.TimeID,
		NewTransaction.LocationID,
		NewTransaction.PaymentMethodID,
		userId,
	).Scan(&transactionID)
	if err != nil {
		return TransactionResult{}, fmt.Errorf("failed to insert transaction: %w", err)
	}

	for _, seat := range NewTransaction.Seats {
		_, err = trx.Exec(ctx, "INSERT INTO transaction_details (id_transaction, seat) VALUES ($1, $2)", transactionID, seat)
		if err != nil {
			return TransactionResult{}, fmt.Errorf("failed to insert seat '%s' into transaction detail: %w", seat, err)
		}
	}
	var result TransactionResult
	query2 := `SELECT t.id, t.name, t.email, t.phone, t.total_payment, t.movie_date, m.title, c.cinema_name, tm.time, l.location, pm.payment_method 
	FROM transactions t 
	JOIN movies m ON m.id = t.id_movie
	JOIN cinemas c ON c.id = t.id_cinema
	JOIN times tm ON tm.id = t.id_time
	JOIN locations l ON l.id = t.id_location
	JOIN payment_methods pm ON pm.id = t.id_payment_method
	WHERE t.id = $1`

	var movieDate time.Time
	var showTime time.Time
	err = trx.QueryRow(ctx, query2, transactionID).Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Phone,
		&result.TotalPayment,
		&movieDate,
		&result.Title,
		&result.CinemaName,
		&showTime,
		&result.Location,
		&result.PaymentMethod,
	)
	if err != nil {
		return TransactionResult{}, err
	}

	result.MovieDate = FormattedDate(movieDate)
	result.Time = FormattedTime(showTime)

	if err := trx.Commit(ctx); err != nil {
		return TransactionResult{}, fmt.Errorf("failed to commit transaction")
	}
	succesed = true

	return result, nil

}
