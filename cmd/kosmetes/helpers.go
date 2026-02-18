package main

import (
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// extract search request from http url query "q", if
// there is no or empty request, return default value
// "status:pending". Return search request and which
// request was used - default (true as second parameter)
// or from url (false as second parameter).
func getSearchRequest(r *http.Request) (string, bool) {
	search := r.URL.Query().Get("q")
	if search == "" {
		return "status:pending", true
	}
	return search, false
}

func (app *application) getTasksBlock(search string) (*TasksBlock, error) {
	tasks, err := app.taskClient.GetTasks(search)
	if err != nil {
		return nil, err
	}

	return &TasksBlock{
		"",
		tasks,
	}, nil
}
