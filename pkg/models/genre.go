package models

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MovieGenre struct {
	ID      int64 `json:"id"`
	MovieId uint  `json:"movieId"`
	GenreId int64 `json:"genreId"`
}
