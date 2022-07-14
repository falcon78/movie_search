package main

import (
	"github.com/golang-migrate/migrate/v4"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/falcon78/movie_search/pkg/repository"
	"github.com/falcon78/movie_search/pkg/utils"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var MigrationDir = "file://../../migrations"
var CsvFile = "../../movies_metadata_new.csv"

func main() {
	db, err := utils.GetDb()
	if err != nil {
		panic(err)
	}

	// Run database migrations
	if m, err := migrate.New(MigrationDir, utils.GetDatabaseStringForMigrate()); err != nil {
		panic(err)
	} else {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			panic(err)
		}
	}

	if strings.ToLower(os.Getenv("INIT_DB")) == "true" {
		if movieInfos, productions, genres, err := utils.InsertCsvToDatabase(CsvFile); err != nil {
			panic(err)
		} else {
			repo := repository.NewRepository(db)
			if err := repo.InsertAllGenres(genres); err != nil {
				panic(err)
			}
			if err := repo.InsertAllProductions(productions); err != nil {
				panic(err)
			}
			if err := repo.InsertMovieInfo(movieInfos); err != nil {
				panic(err)
			}
		}
	}

	app := newApp(db)

	e := echo.New()
	e.Use()
	e.Use(middleware.Gzip())

	// Api Routes
	e.GET("/api/search/movies", app.searchMovies)
	// e.POST("/api/channel/create/:channelName", app.createChannel, basicAuth())
	// e.DELETE("/api/channel/delete/:channelId", app.deleteChannel, basicAuth())

	// Serve static assets for frontend
	e.Static("/assets", "../../static/assets")
	e.File("/*", "../../static/index.html")

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
