package api

type App struct {
	Name   string `json:"name"`
	Repo   string `json:"repo"`
	Port   int    `json:"port"`
	Status string `json:"status"`
}
