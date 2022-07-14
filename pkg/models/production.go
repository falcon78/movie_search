package models

type Production struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MovieProduction struct {
	ID           int64 `json:"id"`
	MovieId      uint  `json:"movieId"`
	ProductionId int64 `json:"productionId"`
}
