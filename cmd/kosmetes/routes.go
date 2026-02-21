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
	mux.HandleFunc("GET /search/{$}", app.getSearch)
	mux.HandleFunc("GET /edit/{uuid}", app.GetTaskEdit)
	mux.HandleFunc("POST /done/{uuid}", app.postDone)
	mux.HandleFunc("POST /undone/{uuid}", app.postUndone)
	mux.HandleFunc("POST /update/{uuid}", app.PostTaskUpdate)

	return mux

}
