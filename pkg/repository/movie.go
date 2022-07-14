package repository

import (
	"fmt"
	"github.com/falcon78/movie_search/pkg/models"
	"gorm.io/gorm"
)

type MovieSearchParam struct {
	Title    string
	PageSize int
	Page     int
	Offset   int
}

func (r *Repo) GetMovies() (*[]models.Movie, error) {
	var movies []models.Movie
	if err := r.DB.Find(&movies).Error; err != nil {
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

func (r *Repo) Paginate(page int, pageSize int) *gorm.DB {
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return r.DB.Offset(offset).Limit(pageSize)
}

func (r *Repo) SearchMovie(param MovieSearchParam) (*[]models.Movie, error) {
	if param.Page == 0 {
		param.Page = 1
	}

	var movies []models.Movie

	err := r.DB.Offset(param.Offset).Limit(param.PageSize).Where("title ILIKE '%?%'", param.Title).Find(&movies).Error
	if err != nil {
		return &movies, err
	}

	return &movies, nil
}
