package main

import (
	"net/http"
)

type MainPage struct {
	Search string
	Blocks []TasksBlock
}

func (app *application) getMainPage(w http.ResponseWriter, r *http.Request) {
	search, defaultSearch := getSearchRequest(r)
	app.logger.Info("getMainPage", "search", search)

	block, err := app.getTasksBlock(search)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	searchQuery := search
	if defaultSearch {
		// don't show default query in search bar
		searchQuery = ""
	}

	err = app.template.ExecuteTemplate(w, "base", &MainPage{
		searchQuery,
		[]TasksBlock{*block},
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) getSearch(w http.ResponseWriter, r *http.Request) {
	search, _ := getSearchRequest(r)
	app.logger.Info("getSearch", "search", search)

	block, err := app.getTasksBlock(search)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = app.template.ExecuteTemplate(w, "block", &MainPage{
		"",
		[]TasksBlock{*block},
	})

	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
