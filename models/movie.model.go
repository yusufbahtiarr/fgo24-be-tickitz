package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"time"

	"github.com/jackc/pgx/v5"
)

type UpcomingMovie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       string    `json:"genre"`
	Rating      float32   `json:"rating"`
}

type Movie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	BackdropURL string    `json:"backdrop_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	Overview    string    `json:"overview"`
	Rating      float32   `json:"rating"`
	IDDirector  int       `json:"id_director"`
}

type Movies []Movie

func FindUpcomingMovies() ([]UpcomingMovie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []UpcomingMovie{}, err
	}
	defer conn.Close()

	query2 := `SELECT m.id, m.poster_url, m.title, m.release_date,
  ( SELECT STRING_AGG(g.genre_name, ', ') 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre, m.rating
FROM movies m;`
	rows, err := conn.Query(context.Background(), query2)
	if err != nil {
		return []UpcomingMovie{}, err
	}

	movies, err := pgx.CollectRows[UpcomingMovie](rows, pgx.RowToStructByName)
	if err != nil {
		return []UpcomingMovie{}, err
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

func FindMovieByID(id int) (Movie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movie{}, err
	}
	defer conn.Close()

	query := "SELECT id, poster_url, backdrop_url, title, release_date, runtime, overview, rating, id_director FROM movies WHERE id = $1"
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return Movie{}, err
	}

	movie, err := pgx.CollectOneRow[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movie{}, err
	}

	return movie, err
}
