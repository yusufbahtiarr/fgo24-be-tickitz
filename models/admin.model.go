package models

import (
	"context"
	"errors"
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
	Genre       []string  `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     string    `json:"runtime"`
}

type DetailCreatedMovie struct {
	ID          int       `json:"id"`
	PosterUrl   string    `json:"poster_url"`
	BackdropUrl string    `json:"backdrop_url"`
	Title       string    `json:"title"`
	Genre       []string  `json:"genre"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     string    `json:"runtime"`
	Director    []string  `json:"director"`
	Cast        []string  `json:"cast"`
	Overview    string    `json:"overview"`
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

func GetAllMoviesCreated(releaseMonth string, limit, offset int) ([]CreatedMovies, int, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return []CreatedMovies{}, 0, err
	}
	defer conn.Close()

	query := `SELECT m.id, m.poster_url, m.title,
  (
    SELECT ARRAY_AGG(g.genre_name) 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre, m.release_date, m.runtime
	FROM movies m 
	WHERE ($1 = '' OR TO_CHAR(m.release_date, 'YYYY-MM') = $1)
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3;`
	rows, err := conn.Query(context.Background(), query, releaseMonth, limit, offset)
	if err != nil {
		return []CreatedMovies{}, 0, err
	}

	movies, err := pgx.CollectRows[CreatedMovies](rows, pgx.RowToStructByName)
	if err != nil {
		return []CreatedMovies{}, 0, err
	}

	Count := `SELECT COUNT(*)	FROM movies WHERE ($1 = '' OR TO_CHAR(release_date, 'YYYY-MM') = $1)`
	var totalMovies int
	err = conn.QueryRow(context.Background(), Count, releaseMonth).Scan(&totalMovies)
	if err != nil {
		return []CreatedMovies{}, 0, err
	}

	return movies, totalMovies, err
}

func GetMovieCreatedByID(id int) (DetailCreatedMovie, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return DetailCreatedMovie{}, err
	}
	defer conn.Close()

	query := `SELECT m.id, m.poster_url, m.backdrop_url, m.title,
  (
    SELECT ARRAY_AGG(g.genre_name) 
    FROM movie_genres mg
    JOIN genres g ON g.id = mg.id_genre
    WHERE mg.id_movie = m.id
  ) AS genre, m.release_date, m.runtime,
	 (
    SELECT ARRAY_AGG(d.director_name) 
    FROM movie_directors md
    JOIN directors d ON d.id = md.id_director
    WHERE md.id_movie = m.id
  ) AS director,
	 (
    SELECT ARRAY_AGG(c.cast_name) 
    FROM movie_casts mc
    JOIN casts c ON c.id = mc.id_cast
    WHERE mc.id_movie = m.id
  ) AS cast, m.overview
 FROM movies m WHERE id = $1`
	rows, err := conn.Query(context.Background(), query, id)
	if err != nil {
		return DetailCreatedMovie{}, err
	}

	movie, err := pgx.CollectOneRow[DetailCreatedMovie](rows, pgx.RowToStructByName)
	if err != nil {
		return DetailCreatedMovie{}, err
	}

	return movie, err
}

func DeleteMovie(id int) error {
	conn, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := `DELETE from movies WHERE id = $1`
	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("movie not found")
	}

	return nil
}

func UpdateMovie(idMovie int, newData dto.UpdateMovieRequest) (dto.UpdateMovieRequest, error) {
	conn, err := db.ConnectDB()
	if err != nil {
		return dto.UpdateMovieRequest{}, err
	}
	defer conn.Close()

	if newData.Title == "" &&
		newData.BackdropUrl == "" &&
		newData.PosterUrl == "" &&
		newData.ReleaseDate == "" &&
		newData.Runtime == 0 &&
		newData.Overview == "" &&
		newData.Rating == 0 &&
		len(newData.Genres) == 0 &&
		len(newData.Casts) == 0 &&
		len(newData.Directors) == 0 {
		return dto.UpdateMovieRequest{}, err
	}

	ctx := context.Background()
	trx, err := conn.Begin(ctx)
	if err != nil {
		return dto.UpdateMovieRequest{}, err
	}
	committed := false
	defer func() {
		if committed {
			_ = trx.Rollback(ctx)
		}
	}()

	oldData := Movie{}
	query := `SELECT title, backdrop_url, poster_url, release_date, runtime, overview, rating FROM movies WHERE id = $1`
	err = trx.QueryRow(ctx, query, idMovie).Scan(
		&oldData.Title,
		&oldData.BackdropURL,
		&oldData.PosterURL,
		&oldData.ReleaseDate,
		&oldData.Runtime,
		&oldData.Overview,
		&oldData.Rating,
	)
	if err != nil {
		return dto.UpdateMovieRequest{}, err
	}

	releaseDate := oldData.ReleaseDate
	if newData.ReleaseDate != "" {
		parsedDate, erra := time.Parse("2006-01-02", newData.ReleaseDate)
		if erra != nil {
			return dto.UpdateMovieRequest{}, erra
		}
		if !parsedDate.Equal(oldData.ReleaseDate) {
			releaseDate = parsedDate
		}
	}
	if newData.Title != "" && newData.Title != oldData.Title {
		oldData.Title = newData.Title
	}
	if newData.BackdropUrl != "" && newData.BackdropUrl != oldData.BackdropURL {
		oldData.BackdropURL = newData.BackdropUrl
	}
	if newData.PosterUrl != "" && newData.PosterUrl != oldData.PosterURL {
		oldData.PosterURL = newData.PosterUrl
	}

	if newData.Runtime > 0 && newData.Runtime != oldData.Runtime {
		oldData.Runtime = newData.Runtime
	}
	if newData.Overview != "" && newData.Overview != oldData.Overview {
		oldData.Overview = newData.Overview
	}
	if newData.Rating > 0 && newData.Rating != oldData.Rating {
		oldData.Rating = newData.Rating
	}

	_, err = trx.Exec(ctx, `
	UPDATE movies 
	SET poster_url = $1, backdrop_url = $2, title = $3, release_date = $4, runtime = $5, overview = $6, rating = $7, updated_at = now()
	WHERE id = $8`,
		oldData.PosterURL,
		oldData.BackdropURL,
		oldData.Title,
		releaseDate,
		oldData.Runtime,
		oldData.Overview,
		oldData.Rating,
		idMovie,
	)
	if err != nil {
		return dto.UpdateMovieRequest{}, err
	}

	if len(newData.Genres) > 0 {
		_, _ = trx.Exec(ctx, `DELETE FROM movie_genres WHERE id_movie = $1`, idMovie)
		for _, idGenre := range newData.Genres {
			if _, err = trx.Exec(ctx, `INSERT INTO movie_genres (id_movie, id_genre) VALUES ($1, $2)`, idMovie, idGenre); err != nil {
				return dto.UpdateMovieRequest{}, err
			}
		}
	}

	if len(newData.Casts) > 0 {
		_, _ = trx.Exec(ctx, `DELETE FROM movie_casts WHERE id_movie = $1`, idMovie)
		for _, idCast := range newData.Casts {
			if _, err = trx.Exec(ctx, `INSERT INTO movie_casts (id_movie, id_cast) VALUES ($1, $2)`, idMovie, idCast); err != nil {
				return dto.UpdateMovieRequest{}, err
			}
		}
	}

	if len(newData.Directors) > 0 {
		_, _ = trx.Exec(ctx, `DELETE FROM movie_directors WHERE id_movie = $1`, idMovie)
		for _, idDirector := range newData.Directors {
			if _, err = trx.Exec(ctx, `INSERT INTO movie_directors (id_movie, id_director) VALUES ($1, $2)`, idMovie, idDirector); err != nil {
				return dto.UpdateMovieRequest{}, err
			}
		}
	}

	if err = trx.Commit(ctx); err != nil {
		return dto.UpdateMovieRequest{}, err
	}
	committed = true

	return newData, nil
}
