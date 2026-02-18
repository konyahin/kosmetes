package main

import (
	"net/http"
)

type MainPage struct {
	Blocks []TasksBlock
}

func (app *application) getMainPage(w http.ResponseWriter, r *http.Request) {
	tasks, err := app.taskClient.GetTasks("status:pending")
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	block := TasksBlock{
		"All tasks",
		tasks,
	}

	err = app.template.ExecuteTemplate(w, "base", &MainPage{[]TasksBlock{block}})
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}
