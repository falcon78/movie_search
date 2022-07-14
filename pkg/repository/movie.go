package repository

import (
	"fmt"
	"github.com/falcon78/movie_search/pkg/models"
)

func (r *Repo) GetMovies() (*[]models.Movie, error) {
	var movies []models.Movie
	if err := r.DB.Find(&movies).Error; err != nil {
		return &movies, err
	}
	return &movies, nil
}

//func (r *Repo) InsertMovies(movieInfos []models.MovieInfo) error {
//	var movies []models.Movie
//	for _, m := range movieInfos {
//		movies = append(movies, m.Movie)
//	}
//
//	i := 0
//	for {
//		upperBound := int(math.Min(float64(len(movies)), float64(i+100)))
//		if err := r.DB.Create(movies[i:upperBound]).Error; err != nil {
//			return err
//		}
//		i += 100
//		if i > len(movies) {
//			break
//		}
//	}
//
//	return nil
//}

func (r *Repo) InsertMovies(movieInfos []models.MovieInfo) error {
	for i, mi := range movieInfos {
		if err := r.DB.Create(&mi.Movie).Error; err != nil {
			return err
		}
		if err := r.InsertGenreRelation(mi); err != nil {
			return err
		}
		if err := r.InsertProductionRelation(mi); err != nil {
			return err
		}
		if i%1000 == 0 {
			fmt.Println(mi.Movie.ID)
		}
	}

	return nil
}
