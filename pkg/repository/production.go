package repository

import (
	"github.com/falcon78/movie_search/pkg/models"
	"math"
)

func (r *Repo) InsertAllProductions(productions []models.Production) error {
	return r.DB.Create(&productions).Error
}

func (r *Repo) InsertProductionRelation(info models.MovieInfo) error {
	if len(info.Productions) == 0 {
		return nil
	}

	var productions []models.MovieProduction
	for _, productionId := range info.Productions {
		productions = append(productions, models.MovieProduction{
			MovieId:      info.Movie.ID,
			ProductionId: productionId,
		})
	}

	return r.DB.Create(&productions).Error
}

func (r *Repo) InsertAllMovieProduction(movieProduction []models.MovieProduction) error {
	i := 0
	for {
		upperBound := int(math.Min(float64(len(movieProduction)), float64(i+100)))
		if err := r.DB.Create(movieProduction[i:upperBound]).Error; err != nil {
			return err
		}
		i += 1000
		if i > len(movieProduction) {
			break
		}
	}
	return nil
}
