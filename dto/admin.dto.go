package dto

import "time"

type CreateMovieRequest struct {
	Title       string    `json:"title" binding:"required"`
	BackdropUrl string    `json:"backdrop_url" binding:"required"`
	PosterUrl   string    `json:"poster_url" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Runtime     int       `json:"runtime" binding:"required"`
	Overview    string    `json:"overview" binding:"required"`
	Rating      float64   `json:"rating" binding:"required"`
	Genres      []int     `json:"genres" binding:"required"`
	Casts       []int     `json:"casts" binding:"required"`
	Directors   []int     `json:"directors" binding:"required"`
}

type UpdateMovieRequest struct {
	Title       string    `json:"title" binding:"required"`
	BackdropUrl string    `json:"backdrop_url" binding:"required"`
	PosterUrl   string    `json:"poster_url" binding:"required"`
	ReleaseDate time.Time `json:"release_date" time_format:"2006-01-02" binding:"required" `
	Runtime     int       `json:"runtime" binding:"required"`
	Overview    string    `json:"overview" binding:"required"`
	Rating      float64   `json:"rating" binding:"required"`
	Genres      []int     `json:"genres" binding:"required"`
	Casts       []int     `json:"casts" binding:"required"`
	Directors   []int     `json:"directors" binding:"required"`
}
