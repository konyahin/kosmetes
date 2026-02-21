package taskwarrior

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

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

func (c *TaskWarriorClient) GetTasks(filter string) ([]model.Task, error) {
	args := strings.Split(filter, " ")
	args = append(args, "export")

	output, err := c.execute(args...)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	if err := json.Unmarshal(output, &tasks); err != nil {
		return nil, fmt.Errorf("can't parse task output: %v", err)
	}

	return tasks, nil
}

func (c *TaskWarriorClient) Done(uuid string) error {
	_, err := c.execute(uuid, "done")
	return err
}

func (c *TaskWarriorClient) Undone(uuid string) error {
	_, err := c.execute(uuid, "modify", "status:pending")
	return err
}

func (c *TaskWarriorClient) UpdateTask(uuid, task string) error {
	args := append([]string{uuid, "modify"}, strings.Split(task, " ")...)
	_, err := c.execute(args...)
	return err
}

func (c *TaskWarriorClient) execute(args ...string) ([]byte, error) {
	var writer bytes.Buffer
	var errWriter bytes.Buffer

	cmd := exec.Command("task", args...)
	cmd.Stdout = &writer
	cmd.Stderr = &errWriter

	if err := cmd.Run(); err != nil {
		return nil, *c.error(err, errWriter.String())
	}

	return writer.Bytes(), nil
}
