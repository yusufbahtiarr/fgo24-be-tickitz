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
	Genre       *string   `json:"genre"`
}

type NowShowingMovie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       *string   `json:"genre"`
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
}

type Movies []Movie

func GetUpcomingMovies() ([]UpcomingMovie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []UpcomingMovie{}, err
	}
	defer conn.Close()

	query2 := `SELECT m.id, m.poster_url, m.title, m.release_date,
  (
    SELECT STRING_AGG(g.genre_name, ', ') 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre
	FROM movies m 
	WHERE m.release_date > CURRENT_DATE;`
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

func GetNowShowingMovies() ([]NowShowingMovie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []NowShowingMovie{}, err
	}
	defer conn.Close()

	query := `SELECT m.id, m.poster_url, m.title, m.release_date,
	(
		SELECT STRING_AGG(g.genre_name, ', ')
		FROM movie_genres mg
		JOIN genres g ON g.id = mg.id_genre
		WHERE mg.id_movie = m.id
	) AS genre, m.rating
		FROM movies m
		WHERE m.release_date BETWEEN CURRENT_DATE - INTERVAL '2 month' AND CURRENT_DATE
		ORDER BY m.release_date DESC, m.rating DESC;`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []NowShowingMovie{}, err
	}

	movies, err := pgx.CollectRows[NowShowingMovie](rows, pgx.RowToStructByName)
	if err != nil {
		return []NowShowingMovie{}, err
	}

	return movies, err
}

func GetListMovies(title string, limit, offset int) (Movies, int, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movies{}, 0, err
	}
	defer conn.Close()

	search := "%" + title + "%"

	query := `
		SELECT id, poster_url, backdrop_url, title, release_date, runtime, overview, rating
		FROM movies
		WHERE title ILIKE $1
		ORDER BY title
		LIMIT $2 OFFSET $3
	`
	rows, err := conn.Query(context.Background(), query, search, limit, offset)
	if err != nil {
		return Movies{}, 0, err
	}
	defer rows.Close()

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movies{}, 0, err
	}

	Count := `SELECT COUNT(*)	FROM movies	WHERE title ILIKE $1	`
	var totalMovies int
	err = conn.QueryRow(context.Background(), Count, search).Scan(&totalMovies)
	if err != nil {
		return Movies{}, 0, err
	}

	return movies, totalMovies, nil
}

func GetMovieByID(id int) (Movie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movie{}, err
	}
	defer conn.Close()

	query := "SELECT id, poster_url, backdrop_url, title, release_date, runtime, overview, rating FROM movies WHERE id = $1"
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return Movie{}, err
	}

	movie, err := pgx.CollectOneRow[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movie{}, err
	}

	return movie, nil
}
