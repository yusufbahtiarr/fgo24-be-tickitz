package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fgo24-be-tickitz/dto"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	IdProfile string `json:"id_profile"`
}

type UserProfile struct {
	Email    string  `json:"email"`
	Fullname *string `json:"fullname"`
	Phone    *string `json:"phone"`
}

type OldUserProfile struct {
	Email     string  `json:"email"`
	Fullname  *string `json:"fullname"`
	Phone     *string `json:"phone"`
	Password  string  `json:"password"`
	IdProfile int     `json:"id_profile"`
}

type Transaction struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	TotalPayment int       `json:"total_payment"`
	MovieDate    time.Time `json:"movie_date"`
	Title        string    `json:"title"`
	Seats        []string  `json:"seats"`
	Cinema       string    `json:"cinema"`
	Time         time.Time `json:"time"`
	Location     string    `json:"location"`
}

func FindUserByEmail(email string) (User, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, email, password, role, id_profile FROM users WHERE email = $1`
	rows, err := conn.Query(context.Background(), query, email)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func FindUserByID(id int) (User, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, email, password, role, id_profile FROM users WHERE id = $1`
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func FindUserProfile(id int) (UserProfile, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return UserProfile{}, err
	}
	defer conn.Close()

	query := `SELECT u.email, p.fullname, p.phone FROM users u 
	join profiles p on p.id = u.id_profile 
	WHERE u.id = $1`
	row, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return UserProfile{}, err
	}

	user, err := pgx.CollectOneRow[UserProfile](row, pgx.RowToStructByName)
	if err != nil {
		return UserProfile{}, err
	}
	return user, err
}

func UpdateProfile(id int, newData dto.UpdateProfileRequest) error {
	conn, errConn := db.ConnectDB()
	if errConn != nil {
		return errConn
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	trx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			trx.Rollback(context.Background())
		}
	}()

	oldData := OldUserProfile{}

	query := `SELECT p.fullname, u.email, p.phone, u.password, u.id_profile FROM users u
 JOIN profiles p ON p.id = u.id_profile 
 WHERE u.id = $1`
	err = conn.QueryRow(context.Background(), query, id).Scan(
		&oldData.Fullname,
		&oldData.Email,
		&oldData.Phone,
		&oldData.Password,
		&oldData.IdProfile,
	)
	if err != nil {
		return fmt.Errorf("failed to get user data: %w", err)
	}

	if newData.NewPassword != "" {
		if newData.ConfirmPassword == "" {
			err = fmt.Errorf("confirmation password is required when setting a new password")
			return err
		}
		if len(newData.NewPassword) < 8 {
			err = fmt.Errorf("password must be at least 8 characters long")
			return err
		}
		if newData.NewPassword != newData.ConfirmPassword {
			err = fmt.Errorf("new password and confirmation password must match")
			return err
		}
		hashedPassword, errHash := bcrypt.GenerateFromPassword(
			[]byte(newData.NewPassword),
			bcrypt.DefaultCost,
		)
		if errHash != nil {
			err = fmt.Errorf("failed to hash password: %w", errHash)
			return err
		}
		oldData.Password = string(hashedPassword)
	}

	if newData.Fullname != "" && (oldData.Fullname == nil || newData.Fullname != *oldData.Fullname) {
		oldData.Fullname = &newData.Fullname
	}
	if newData.Email != "" && newData.Email != oldData.Email {
		oldData.Email = newData.Email
	}
	if newData.Phone != "" && (oldData.Phone == nil || newData.Phone != *oldData.Phone) {
		oldData.Phone = &newData.Phone
	}

	_, err = trx.Exec(context.Background(),
		`UPDATE users SET email = $1, password = $2 WHERE id = $3 
 `,
		oldData.Email, oldData.Password, id,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	_, err = trx.Exec(context.Background(),
		`UPDATE profiles SET fullname = $1, phone = $2 WHERE id = $3`,
		oldData.Fullname, oldData.Phone, oldData.IdProfile,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	if err = trx.Commit(context.Background()); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return err
}

func GetTransactionHistory(id int) ([]Transaction, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []Transaction{}, err
	}
	defer conn.Close()

	query := `
	SELECT t.id, t.name, t.email, t.phone, t.total_payment, t.movie_date, m.title, 
  (
    SELECT ARRAY_AGG(td.seat) 
    FROM transaction_details td
    WHERE td.id_transaction = t.id
  ) AS seats, 
  	c.cinema_name AS cinema, tm.time, l.location 
	FROM transactions t
	JOIN movies m ON m.id = t.id_movie
	JOIN cinemas c ON c.id = t.id_cinema
	JOIN times tm ON tm.id = t.id_time
	JOIN locations l ON l.id = t.id_location
	WHERE t.id_user = $1
	ORDER BY t.created_at DESC;
	`
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return []Transaction{}, err
	}
	defer rows.Close()

	history, err := pgx.CollectRows[Transaction](rows, pgx.RowToStructByName)
	if err != nil {
		return []Transaction{}, err
	}

	return history, err
}
