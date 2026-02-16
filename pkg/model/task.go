package model

import "strings"

type TaskStatus string

const (
	Completed TaskStatus = "completed"
	Pending   TaskStatus = "pending"
)

type Task struct {
	Description string     `json:"description"`
	Project     string     `json:"project"`
	Status      TaskStatus `json:"status"`
	Uuid        string     `json:"uuid"`
}

func (t *Task) String() string {
	var buf strings.Builder
	switch t.Status {
	case Completed:
		buf.WriteString("x ")
	case Pending:
		buf.WriteString("- ")
	default:
		buf.WriteString("  ")
	}

	buf.WriteString(t.Description)

	return buf.String()
}

func (t *Task) IsCompleted() bool {
	return t.Status == Completed
}
