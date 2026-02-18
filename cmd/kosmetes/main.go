package main

import (
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/konyahin/kosmetes/pkg/model"
	"github.com/konyahin/kosmetes/pkg/taskwarrior"
)

type application struct {
	taskClient *taskwarrior.TaskWarriorClient
	template   *template.Template
	logger     *slog.Logger
}

type TasksBlock struct {
	Title string
	Tasks []model.Task
}

func main() {
	app := &application{
		&taskwarrior.TaskWarriorClient{},
		template.Must(template.ParseGlob("./internal/web/templates/*.html")),
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	log.Println("Starting HTTP server on http://localhost:8100")
	http.ListenAndServe(":8100", app.routes())
}
