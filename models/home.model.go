package models

import (
	"context"
	"fgo24-be-tickitz/db"

	"github.com/jackc/pgx/v5"
)

func FindUpcomingMovies() (Movies, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movies{}, err
	}
	defer conn.Close()

	query := "SELECT id, poster_url, backdrop_url, title, release_date, runtime, overview, rating, id_director FROM movies ORDER BY release_date DESC"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return Movies{}, err
	}

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movies{}, err
	}

	return movies, err
}

func FindListMovies(title string) (Movies, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movies{}, err
	}
	defer conn.Close()

	search := "%" + title + "%"

	query := "SELECT id, poster_url, backdrop_url, title, release_date, runtime, overview, rating, id_director FROM movies WHERE title ILIKE $1"
	rows, err := conn.Query(context.Background(), query, search)
	if err != nil {
		return Movies{}, err
	}

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movies{}, err
	}

	return movies, err
}
