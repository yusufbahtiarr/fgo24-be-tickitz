package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fgo24-be-tickitz/dto"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
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

func ResetPassword(id int, newData dto.ResetPasswordRequest) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}
	defer conn.Close()

	if newData.Password == "" && newData.ConfirmPassword == "" {
		return fmt.Errorf("input password cannot be empty")
	}

	if newData.Password != newData.ConfirmPassword {
		return fmt.Errorf("password and confirm password do not match")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newData.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set password = $1, updated_at = now() where id=$2`, hashedPassword, id)

	return err
}

func VerifyResetPasswordToken(tokenString string) (int, error) {
	secretKey := os.Getenv("APP_SECRET")
	if secretKey == "" {
		return 0, fmt.Errorf("missing APP_SECRET")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["purpose"] != "reset_password" {
		return 0, fmt.Errorf("invalid token purpose")
	}

	userIDStr, ok := claims["userId"].(string)
	if !ok {
		return 0, fmt.Errorf("invalid user ID in token")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, fmt.Errorf("invalid user ID format")
	}

	return userID, nil
}
