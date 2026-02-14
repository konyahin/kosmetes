package taskwarrior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/konyahin/kosmetes/pkg/model"
)

type TaskWarriorClient struct {
	bin string
}

type TaskWarriorError struct {
	Client *TaskWarriorClient
	Err    error
	Stderr string
}

func (err TaskWarriorError) Error() string {
	return fmt.Sprintf("%s command failed: %v\nstderr:%v\n",
		err.Client.getBin(), err.Err, err.Stderr)
}

func NewTaskWarriorClient(bin string) *TaskWarriorClient {
	if bin == "" {
		bin = "task"
	}
	return &TaskWarriorClient{
		bin,
	}
}

func (c *TaskWarriorClient) getBin() string {
	return c.bin
}

func (c *TaskWarriorClient) error(err error, stderr string) *TaskWarriorError {
	return &TaskWarriorError{c, err, stderr}
}

func (c *TaskWarriorClient) GetTasks(filter model.Filter) (error, []model.Task) {
	var writer bytes.Buffer
	var errWriter bytes.Buffer
	cmd := exec.Command("task", filter.Content, "export")
	cmd.Stdout = &writer
	cmd.Stderr = &errWriter

	if err := cmd.Run(); err != nil {
		return *c.error(err, errWriter.String()), nil
	}

	var tasks []model.Task
	if err := json.Unmarshal(writer.Bytes(), &tasks); err != nil {
		return fmt.Errorf("can't parse %s output: %v", c.getBin(), err), nil
	}

	return nil, tasks
}
