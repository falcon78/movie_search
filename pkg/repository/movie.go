package repository

import (
	"fmt"
	"github.com/falcon78/movie_search/pkg/models"
)

type MovieSearchParam struct {
	SearchBy   string
	SearchText string
	PageSize   int
	Page       int
}

type MovieSearchResult struct {
	Movies       []models.Movie `json:"movies"`
	Page         int            `json:"page"`
	PageSize     int            `json:"pageSize"`
	TotalResults int64          `json:"totalResults"`
}

func (r *Repo) GetMovie(id int) (*models.Movie, error) {
	var movies models.Movie
	if err := r.DB.Where("id = ?", id).First(&movies).Error; err != nil {
		return &movies, err
	}
	return &movies, nil
}

func (r *Repo) InsertMovieInfo(movieInfos []models.MovieInfo) error {
	for i, mi := range movieInfos {
		if err := r.DB.Create(&mi.Movie).Error; err != nil {
			return err
		}
		if err := r.InsertGenreRelations(mi); err != nil {
			return err
		}
		if err := r.InsertProductionRelations(mi); err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Println(mi.Movie.ID)
		}
	}

	return nil
}

func (r *Repo) SearchMovie(param MovieSearchParam) (*MovieSearchResult, error) {
	if param.Page == 0 {
		param.Page = 1
	}

	result := MovieSearchResult{
		Movies:       []models.Movie{},
		Page:         param.Page,
		PageSize:     param.PageSize,
		TotalResults: 0,
	}

	offset := (param.Page - 1) * param.PageSize

	if param.SearchBy == "production" {
		if err := r.DB.Raw(
			"select count(*) from movie_production_view where name ILIKE ?",
			fmt.Sprintf("%%%s%%", param.SearchText),
		).Scan(&result.TotalResults).Error; err != nil {
			return &result, err
		}

		if err := r.DB.Raw(
			"select * from movie_production_view where name ILIKE ? limit ? offset ?",
			fmt.Sprintf("%%%s%%", param.SearchText),
			param.PageSize,
			offset,
		).Scan(&result.Movies).Error; err != nil {
			return &result, err
		}

		return &result, nil
	} else if param.SearchBy == "genre" {
		if err := r.DB.Raw(
			"select count(*) from movie_genre_view where name ILIKE ?",
			fmt.Sprintf("%%%s%%", param.SearchText),
		).Scan(&result.TotalResults).Error; err != nil {
			return &result, err
		}

		if err := r.DB.Raw(
			"select * from movie_genre_view where name ILIKE ? limit ? offset ?",
			fmt.Sprintf("%%%s%%", param.SearchText),
			param.PageSize,
			offset,
		).Scan(&result.Movies).Error; err != nil {
			return &result, err
		}

		return &result, nil
	} else {
		if err := r.DB.Model(&models.Movie{}).
			Where("title ILIKE ?", fmt.Sprintf("%%%s%%", param.SearchText)).
			Count(&result.TotalResults).Error; err != nil {
			return nil, err
		}

		if err := r.DB.Offset(offset).
			Limit(param.PageSize).
			Where("title ILIKE ?", fmt.Sprintf("%%%s%%", param.SearchText)).
			Find(&result.Movies).Error; err != nil {
			return &result, err
		}

		return &result, nil
	}
}
