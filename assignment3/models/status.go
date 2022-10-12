package models

type StatusWeather struct {
	Status struct {
		Water int `json:"Water"`
		Wind  int `json:"Wind"`
	} `json:"status"`
}
