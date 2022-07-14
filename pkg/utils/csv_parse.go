package utils

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/falcon78/movie_search/pkg/models"
)

type genres struct {
	Genres []models.Genre `json:"genres"`
}
type productions struct {
	Productions []models.Production `json:"productions"`
}

func InsertCsvToDatabase(path string) ([]models.MovieInfo, []models.Production, []models.Genre, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	if _, err := csvReader.Read(); err != nil {
		return nil, nil, nil, err
	}

	var movieInfos []models.MovieInfo
	allGenres := make(map[int64]models.Genre)
	allProductions := make(map[int64]models.Production)

	for {
		if record, err := csvReader.Read(); err != nil && errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			fmt.Println(err)
		} else {
			var gnrs genres
			var pdns productions

			if err := json.Unmarshal([]byte(record[3]), &gnrs); err != nil {
				return nil, nil, nil, err
			}
			if err := json.Unmarshal([]byte(record[12]), &pdns); err != nil {
				return nil, nil, nil, err
			}

			movie := models.Movie{
				Adult:       Str2Bool(record[0]),
				Budget:      Str2Int(record[2]),
				Homepage:    record[4],
				ImdbId:      record[6],
				Overview:    record[9],
				Popularity:  Str2Float(record[10]),
				Poster:      record[11],
				ReleaseDate: Str2Date(record[14]),
				Revenue:     Str2Int(record[15]),
				Runtime:     Str2Float(record[16]),
				Status:      record[18],
				Title:       record[20],
				VoteAverage: Str2Float(record[22]),
				VoteCount:   Str2Int(record[23]),
			}

			movieInfo := models.MovieInfo{
				Movie:       movie,
				Genres:      []int64{},
				Productions: []int64{},
			}

			for _, g := range gnrs.Genres {
				movieInfo.Genres = append(movieInfo.Genres, g.ID)
				allGenres[g.ID] = g
			}
			for _, p := range pdns.Productions {
				movieInfo.Productions = append(movieInfo.Productions, p.ID)
				allProductions[p.ID] = p
			}

			movieInfos = append(movieInfos, movieInfo)
		}
	}

	var productions []models.Production
	var genres []models.Genre

	for _, p := range allProductions {
		productions = append(productions, p)
	}
	for _, g := range allGenres {
		genres = append(genres, g)
	}

	return movieInfos, productions, genres, nil
}
