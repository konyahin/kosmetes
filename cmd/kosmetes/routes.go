package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./internal/web/static"))))

	mux.HandleFunc("GET /{$}", app.getMainPage)

	return mux

}
