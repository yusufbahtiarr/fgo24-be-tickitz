package models

import (
	"context"
	"fgo24-be-tickitz/db"
	"fgo24-be-tickitz/dto"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type CreatedMovies struct {
	ID          int       `json:"id"`
	PosterUrl   string    `json:"poster_url"`
	Title       string    `json:"title"`
	Genre       string    `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     string    `json:"runtime"`
}

func CreateMovie(movie dto.CreateMovieRequest) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx := context.Background()

	trx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer trx.Rollback(ctx)

	query := `
		INSERT INTO movies (title, backdrop_url, poster_url, release_date, runtime, overview, rating)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	var idMovie int
	err = conn.QueryRow(ctx, query,
		movie.Title,
		movie.BackdropUrl,
		movie.PosterUrl,
		movie.ReleaseDate,
		movie.Runtime,
		movie.Overview,
		movie.Rating,
	).Scan(&idMovie)
	if err != nil {
		return fmt.Errorf("failed to insert movie %w", err)
	}

	for _, idGenre := range movie.Genres {
		_, err := trx.Exec(ctx, "INSERT INTO movie_genres (id_movie, id_genre) VALUES ($1, $2)", idMovie, idGenre)
		if err != nil {
			return fmt.Errorf("failed to insert genre")
		}
	}

	for _, idCast := range movie.Casts {
		_, err := trx.Exec(ctx, "INSERT INTO movie_casts (id_movie, id_cast) VALUES ($1, $2)", idMovie, idCast)
		if err != nil {
			return fmt.Errorf("failed to insert cast")
		}
	}

	for _, idDirector := range movie.Directors {
		_, err := trx.Exec(ctx, "INSERT INTO movie_directors (id_movie, id_director) VALUES ($1, $2)", idMovie, idDirector)
		if err != nil {
			return fmt.Errorf("failed to insert director")
		}
	}

	if err := trx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction")
	}

	return nil
}

func GetAllMoviesCreated() ([]CreatedMovies, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []CreatedMovies{}, err
	}
	defer conn.Close()

	query := `SELECT m.id, m.poster_url, m.title,
  (
    SELECT STRING_AGG(g.genre_name, ', ') 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre, m.release_date, m.runtime
	FROM movies m 
	ORDER BY created_at DESC;`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []CreatedMovies{}, err
	}

	movies, err := pgx.CollectRows[CreatedMovies](rows, pgx.RowToStructByName)
	if err != nil {
		return []CreatedMovies{}, err
	}

	return movies, err
}
