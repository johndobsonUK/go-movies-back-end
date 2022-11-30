package main

import (
	"net/http"
)

var p struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	p.Status = "active"
	p.Message = "Go Movies up and running"
	p.Version = "1.0.0"

	_ = app.writeJSON(w, http.StatusOK, p)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}
