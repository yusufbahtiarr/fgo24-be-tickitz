package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type UpcomingMovie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       []string  `json:"genre"`
}

type NowShowingMovie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       []string  `json:"genre"`
	Rating      float64   `json:"rating"`
}

type Movie struct {
	ID          int       `json:"id"`
	PosterURL   string    `json:"poster_url"`
	BackdropURL string    `json:"backdrop_url"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Genre       []string  `json:"genre"`
	Cast        []string  `json:"cast"`
	Director    []string  `json:"director"`
	Runtime     int       `json:"runtime"`
	Overview    string    `json:"overview"`
	Rating      float64   `json:"rating"`
}

type Genre struct {
	ID        int    `json:"id"`
	GenreName string `json:"genre_name"`
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
    SELECT ARRAY_AGG(g.genre_name) 
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
		SELECT ARRAY_AGG(g.genre_name)
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

func GetListMovies(title, genre, sort string, limit, offset int) (Movies, int, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movies{}, 0, fmt.Errorf("failed to connect to database")
	}
	defer conn.Close()

	search := "%" + title + "%"

	var sorting string
	switch sort {
	case "title_asc":
		sorting = "ORDER BY m.title ASC"
	case "title_desc":
		sorting = "ORDER BY m.title DESC"
	case "popular":
		sorting = "ORDER BY m.rating DESC"
	case "latest":
		sorting = "ORDER BY m.release_date DESC"
	default:
		sorting = ""
	}

	query := fmt.Sprintf(`
		SELECT m.id, m.poster_url, m.backdrop_url, m.title, m.release_date, 
		(
		SELECT ARRAY_AGG(g.genre_name)
		FROM movie_genres mg
		JOIN genres g ON g.id = mg.id_genre
		WHERE mg.id_movie = m.id
	) AS genre, 
	(
		SELECT ARRAY_AGG(c.cast_name)
		FROM movie_casts mc
		JOIN casts c ON c.id = mc.id_cast
		WHERE mc.id_movie = m.id
	) AS cast,
	 (
		SELECT ARRAY_AGG(d.director_name)
		FROM movie_directors md
		JOIN directors d ON d.id = md.id_director
		WHERE md.id_movie = m.id
	) AS director,  	 
	m.runtime, m.overview, m.rating
		FROM movies m
		WHERE m.title ILIKE $1
		AND (
		$2 = '' OR
		EXISTS (
				SELECT 1
				FROM movie_genres mg
				JOIN genres g on g.id = mg.id_genre
				WHERE mg.id_movie = m.id
				AND g.genre_name ILIKE $2
			)
		)
		%s
		LIMIT $3 OFFSET $4`, sorting)
	rows, err := conn.Query(context.Background(), query, search, genre, limit, offset)
	if err != nil {
		return Movies{}, 0, fmt.Errorf("failed to execute movie list query")
	}
	defer rows.Close()

	movies, err := pgx.CollectRows[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movies{}, 0, fmt.Errorf("failed to collect movie rows")
	}

	Count := `SELECT COUNT(*)	FROM movies m	WHERE title ILIKE $1 AND (
		$2 = '' OR
		EXISTS (
				SELECT 1
				FROM movie_genres mg
				JOIN genres g on g.id = mg.id_genre
				WHERE mg.id_movie = m.id
				AND g.genre_name ILIKE $2
			)
		) `
	var totalMovies int
	err = conn.QueryRow(context.Background(), Count, search, genre).Scan(&totalMovies)
	if err != nil {
		return Movies{}, 0, fmt.Errorf("failed to get total movie count")
	}

	return movies, totalMovies, nil
}

func GetMovieByID(id int) (Movie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return Movie{}, err
	}
	defer conn.Close()

	query := `SELECT m.id, m.poster_url, m.backdrop_url, m.title, m.release_date, 
		(
		SELECT ARRAY_AGG(g.genre_name)
		FROM movie_genres mg
		JOIN genres g ON g.id = mg.id_genre
		WHERE mg.id_movie = m.id
	) AS genre, 
	(
		SELECT ARRAY_AGG(c.cast_name)
		FROM movie_casts mc
		JOIN casts c ON c.id = mc.id_cast
		WHERE mc.id_movie = m.id
	) AS cast,
	 (
		SELECT ARRAY_AGG(d.director_name)
		FROM movie_directors md
		JOIN directors d ON d.id = md.id_director
		WHERE md.id_movie = m.id
	) AS director,  
	m.runtime, m.overview, m.rating
		FROM movies m WHERE m.id = $1`

	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return Movie{}, fmt.Errorf("failed query   %w", err)
	}
	movie, err := pgx.CollectOneRow[Movie](rows, pgx.RowToStructByName)
	if err != nil {
		return Movie{}, fmt.Errorf("failed query rows  %w", err)
	}

	return movie, nil
}

func GetGenres() ([]Genre, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []Genre{}, err
	}
	defer conn.Close()

	query := `SELECT id, genre_name FROM genres`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []Genre{}, err
	}

	genres, err := pgx.CollectRows[Genre](rows, pgx.RowToStructByName)
	if err != nil {
		return []Genre{}, err
	}

	return genres, err
}
