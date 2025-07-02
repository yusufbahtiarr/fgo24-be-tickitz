package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fgo24-be-tickitz/dto"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(reqUser dto.RegisterUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}
	conn, err := db.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}
	defer conn.Close()

	ctx := context.Background()

	trx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}

	defer func() {
		if r := recover(); r != nil {
			_ = trx.Rollback(ctx)
			panic(r)
		} else if err != nil {
			_ = trx.Rollback(ctx)
		}
	}()

	queryProfile := `INSERT INTO profile (created_at, updated_at) VALUES (now(), now()) RETURNING id`
	var profileID int
	err = trx.QueryRow(ctx, queryProfile).Scan(&profileID)
	if err != nil {
		return fmt.Errorf("failed to insert profile into profile table")
	}

	queryUser := `INSERT INTO users (email, password, id_profile) VALUES ($1, $2, $3) RETURNING id`
	var userID int
	err = trx.QueryRow(ctx, queryUser, reqUser.Email, hashedPassword, profileID).Scan(&userID)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return fmt.Errorf("email '%s' is already registered", reqUser.Email)
		}
		return fmt.Errorf("failed to insert user into users table")
	}

	err = trx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func CheckUserExistsByEmail(email string) (bool, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return false, fmt.Errorf("failed to connect to database")
	}
	defer conn.Close()

	query := `SELECT id FROM users WHERE email = $1`

	var userID int
	err = conn.QueryRow(context.Background(), query, email).Scan(&userID)

	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}

		return false, fmt.Errorf("failed to check user existence")
	}

	return true, nil
}
