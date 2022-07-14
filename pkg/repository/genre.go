package repository

import (
	"github.com/falcon78/movie_search/pkg/models"
	"math"
)

func (r *Repo) InsertAllGenres(genres []models.Genre) error {
	return r.DB.Create(&genres).Error
}

func (r *Repo) InsertGenreRelation(info models.MovieInfo) error {
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

func (r *Repo) InsertAllGenreRelations(movieGenre []models.MovieGenre) error {
	i := 0
	for {
		upperBound := int(math.Min(float64(len(movieGenre)), float64(i+100)))
		if err := r.DB.Create(movieGenre[i:upperBound]).Error; err != nil {
			return err
		}
		i += 1000
		if i > len(movieGenre) {
			break
		}
	}
	return nil
}
