package dto

type CreateMovieRequest struct {
	Title       string  `form:"title" binding:"required"`
	ReleaseDate string  `form:"release_date" binding:"required"`
	Runtime     int     `form:"runtime" binding:"required"`
	Overview    string  `form:"overview" binding:"required"`
	Rating      float64 `form:"rating" binding:"required"`

	GenresRaw    string `form:"genres" binding:"required"`
	CastsRaw     string `form:"casts" binding:"required"`
	DirectorsRaw string `form:"directors" binding:"required"`

	Genres    []int `form:"-"`
	Casts     []int `form:"-"`
	Directors []int `form:"-"`
}

type UpdateMovieRequest struct {
	Title       string  `form:"title,omitempty"`
	BackdropUrl string  `form:"backdrop_url,omitempty"`
	PosterUrl   string  `form:"poster_url,omitempty"`
	ReleaseDate string  `form:"release_date,omitempty"`
	Runtime     int     `form:"runtime,omitempty"`
	Overview    string  `form:"overview,omitempty"`
	Rating      float64 `form:"rating,omitempty"`
	Genres      []int   `form:"genres,omitempty"`
	Casts       []int   `form:"casts,omitempty"`
	Directors   []int   `form:"directors,omitempty"`
}
