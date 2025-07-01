package models

import "time"

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
