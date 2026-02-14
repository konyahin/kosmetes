package taskwarrior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/konyahin/kosmetes/pkg/model"
)

type TaskWarriorClient struct {
}

type TaskWarriorError struct {
	Client *TaskWarriorClient
	Err    error
	Stderr string
}

func (err TaskWarriorError) Error() string {
	return fmt.Sprintf("task command failed: %v\nstderr:%v\n", err.Err, err.Stderr)
}

func (c *TaskWarriorClient) error(err error, stderr string) *TaskWarriorError {
	return &TaskWarriorError{c, err, stderr}
}

func (c *TaskWarriorClient) GetTasks(filter model.Filter) ([]model.Task, error) {
	var writer bytes.Buffer
	var errWriter bytes.Buffer
	cmd := exec.Command("task", filter.Content, "export")
	cmd.Stdout = &writer
	cmd.Stderr = &errWriter

	if err := cmd.Run(); err != nil {
		return nil, *c.error(err, errWriter.String())
	}

	var tasks []model.Task
	if err := json.Unmarshal(writer.Bytes(), &tasks); err != nil {
		return nil, fmt.Errorf("can't parse task output: %v", err)
	}

	return tasks, nil
}
