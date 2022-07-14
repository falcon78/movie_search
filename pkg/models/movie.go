package models

import "time"

type Movie struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Adult       bool      `json:"adult"`
	Budget      int64     `json:"budget"`
	Homepage    string    `json:"homepage"`
	ImdbId      string    `json:"imdbId"`
	Overview    string    `json:"overview"`
	Popularity  float64   `json:"popularity"`
	Poster      string    `json:"poster"`
	ReleaseDate time.Time `json:"releaseDate"`
	Revenue     int64     `json:"revenue"`
	Runtime     float64   `json:"runtime"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	VoteAverage float64   `json:"voteAverage"`
	VoteCount   int64     `json:"voteCount"`
}

type MovieInfo struct {
	Movie
	Productions []int64
	Genres      []int64
}
