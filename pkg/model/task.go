package model

import "strings"

type TaskStatus string

const (
	Completed TaskStatus = "completed"
	Pending   TaskStatus = "pending"
)

type Task struct {
	Uuid        string     `json:"uuid"`
	Description string     `json:"description"`
	Project     string     `json:"project"`
	Status      TaskStatus `json:"status"`
	Tags        []string   `json:"tags"`
}

func (t *Task) FullText() string {
	var buf strings.Builder
	if t.Project != "" {
		buf.WriteString("project:")
		buf.WriteString(t.Project)
		buf.WriteRune(' ')
	}

	for _, tag := range t.Tags {
		buf.WriteByte('+')
		buf.WriteString(tag)
		buf.WriteRune(' ')
	}

	buf.WriteString(t.Description)

	return buf.String()
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
