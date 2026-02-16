package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/konyahin/kosmetes/pkg/model"
	"github.com/konyahin/kosmetes/pkg/taskwarrior"
)

type TasksBlock struct {
	Title string
	Tasks []model.Task
}

type MainPage struct {
	Blocks []TasksBlock
}

func main() {
	var taskClient taskwarrior.TaskWarriorClient

	filters := []model.Filter{
		{
			Name:    "Monthly plan",
			Content: "project:2026.jan",
		},
		{
			Name:    "Weekly plan",
			Content: "+week",
		},
		{
			Name:    "Kosmetes Development",
			Content: "project:kosmetes",
		},
		{
			Name:    "Test category",
			Content: "+test",
		},
	}

	blocks := make([]TasksBlock, 0, len(filters))
	for _, filter := range filters {
		var block TasksBlock
		block.Title = filter.Name

		tasks, err := taskClient.GetTasks(filter)
		if err != nil {
			log.Fatalf("%v", err)
		}

		block.Tasks = tasks
		blocks = append(blocks, block)
	}

	tmpl := template.Must(template.ParseGlob("./internal/web/templates/*.html"))

	mux := http.NewServeMux()
	
	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./internal/web/static"))))
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", &MainPage{blocks})
		if err != nil {
			log.Printf("can't execute template: %v", err)
		}
	})

	log.Println("Starting HTTP server on http://localhost:8000")
	http.ListenAndServe(":8000", mux)
}
