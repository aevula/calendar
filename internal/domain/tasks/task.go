package tasks

import (
	"time"
)

type TaskID int
type TaskName string
type QueueName string
type TaskPayload map[string]any
type TaskStatus string

const (
	TaskStatusPending   TaskStatus = "pending"
	TaskStatusEnqueued  TaskStatus = "enqueued"
	TaskStatusRunning   TaskStatus = "running"
	TaskStatusSucceeded TaskStatus = "succeeded"
	TaskStatusFailed    TaskStatus = "failed"
)

type Task struct {
	ID        TaskID
	Name      TaskName
	Queue     QueueName
	Payload   TaskPayload
	Status    TaskStatus
	Attempt   int
	StartAt   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
