package main

import (
	"github.com/falcon78/movie_search/pkg/models"
	"github.com/falcon78/movie_search/pkg/repository"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (a *app) searchMovies(c echo.Context) error {
	page, err := strconv.ParseInt(c.QueryParam("page"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "could not parse 'page' query params")
	}

	pageSize, err := strconv.ParseInt(c.QueryParam("pageSize"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "could not parse 'pageSize' query params")
	}

	searchBy := strings.TrimSpace(c.QueryParam("searchBy"))
	if len(searchBy) == 0 {
		searchBy = "title"
	}

	searchText := strings.TrimSpace(c.QueryParam("searchText"))
	if len(searchText) == 0 {
		return c.String(http.StatusBadRequest, "please specify searchText in query param")
	}

	params := repository.MovieSearchParam{
		SearchBy:   searchBy,
		SearchText: searchText,
		PageSize:   int(pageSize),
		Page:       int(page),
	}

	repo := repository.NewRepository(a.db)
	if movies, err := repo.SearchMovie(params); err != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	} else {
		return c.JSON(http.StatusOK, &movies)
	}
}

type MovieAdditionalInfo struct {
	Genres      []models.Genre      `json:"genres"`
	Productions []models.Production `json:"productions"`
}

func (a *app) GetMovieInfo(c echo.Context) error {
	movieId, err := strconv.ParseInt(c.Param("movieId"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad movie id")
	}

	repo := repository.NewRepository(a.db)
	movie, err := repo.GetMovie(int(movieId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "errored when retrieving movie from database")
	}

	return c.JSON(http.StatusOK, movie)
}

func (a *app) GetMovieAdditionalInfo(c echo.Context) error {
	movieId, err := strconv.ParseInt(c.Param("movieId"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad movie id")
	}

	repo := repository.NewRepository(a.db)
	genres, err := repo.GetAllGenresForMovie(uint(movieId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "error while retrieving genre data")
	}

	productions, err := repo.GetAllProductionsForMovie(uint(movieId))
	if err != nil {
		return c.String(http.StatusInternalServerError, "error while retrieving production data")
	}

	return c.JSON(http.StatusOK, &MovieAdditionalInfo{
		Genres:      genres,
		Productions: productions,
	})
}
