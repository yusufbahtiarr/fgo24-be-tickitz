package dto

type CreateMovieRequest struct {
	Title       string  `json:"title" binding:"required"`
	BackdropUrl string  `json:"backdrop_url" binding:"required"`
	PosterUrl   string  `json:"poster_url" binding:"required"`
	ReleaseDate string  `json:"release_date" binding:"required"`
	Runtime     int     `json:"runtime" binding:"required"`
	Overview    string  `json:"overview" binding:"required"`
	Rating      float64 `json:"rating" binding:"required"`
	Genres      []int   `json:"genres" binding:"required"`
	Casts       []int   `json:"casts" binding:"required"`
	Directors   []int   `json:"directors" binding:"required"`
}

type UpdateMovieRequest struct {
	Title       string  `json:"title,omitempty"`
	BackdropUrl string  `json:"backdrop_url,omitempty"`
	PosterUrl   string  `json:"poster_url,omitempty"`
	ReleaseDate string  `json:"release_date,omitempty"`
	Runtime     int     `json:"runtime,omitempty"`
	Overview    string  `json:"overview,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	Genres      []int   `json:"genres,omitempty"`
	Casts       []int   `json:"casts,omitempty"`
	Directors   []int   `json:"directors,omitempty"`
}
