package models

import (
	"context"
	"fgo24-be-tickitz/db"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func FindUserByEmail(email string) (User, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, email, password, role FROM users WHERE email = $1`
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

	query := `SELECT id, email, password, role FROM users WHERE id = $1`
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
