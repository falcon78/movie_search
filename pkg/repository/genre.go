package repository

import (
	"github.com/falcon78/movie_search/pkg/models"
)

func (r *Repo) InsertAllGenres(genres []models.Genre) error {
	return r.DB.Create(&genres).Error
}

func (r *Repo) InsertGenreRelations(info models.MovieInfo) error {
	if len(info.Genres) == 0 {
		return nil
	}

	var genres []models.MovieGenre
	for _, genreId := range info.Genres {
		genres = append(genres, models.MovieGenre{
			MovieId: info.Movie.ID,
			GenreId: genreId,
		})
	}

	return r.DB.Create(&genres).Error
}
