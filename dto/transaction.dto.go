package dto

type CreateTransactionRequest struct {
	MovieID         int      `json:"id_movie"`
	MovieDate       string   `json:"movie_date"`
	TimeID          int      `json:"id_time"`
	LocationID      int      `json:"id_location"`
	CinemaID        int      `json:"id_cinema"`
	PaymentMethodID int      `json:"id_payment_method"`
	TotalPayment    int      `json:"total_payment"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Seats           []string `json:"seats"`
}
