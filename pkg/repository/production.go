package repository

import (
	"github.com/falcon78/movie_search/pkg/models"
)

func (r *Repo) InsertAllProductions(productions []models.Production) error {
	return r.DB.Create(&productions).Error
}

func (r *Repo) InsertProductionRelations(info models.MovieInfo) error {
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
