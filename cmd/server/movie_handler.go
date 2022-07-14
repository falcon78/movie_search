package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (a *app) searchMovies(c echo.Context) error {
	return c.String(http.StatusOK, "hey")
}
