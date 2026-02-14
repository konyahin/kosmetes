package main

import (
	"fmt"
	"log"

	"github.com/konyahin/kosmetes/pkg/model"
	"github.com/konyahin/kosmetes/pkg/taskwarrior"
)

func main() {
	var taskClient taskwarrior.TaskWarriorClient

	filters := []model.Filter{
		{
			Name:    "планы на январь",
			Content: "project:2026.jan",
		},
		{
			Name:    "планы на неделю",
			Content: "+week",
		},

		{
			Name:    "AI агент для стажеров",
			Content: "project:ai-intern",
		},
		{
			Name:    "онбординг для тимлидов",
			Content: "project:onboarding",
		},
	}

	for _, filter := range filters {
		fmt.Printf("# %s\n", filter.Name)

		tasks, err := taskClient.GetTasks(filter)
		if err != nil {
			log.Fatalf("%v", err)
		}

		for _, task := range tasks {
			fmt.Println(task.String())
		}

		fmt.Println()
	}
}
