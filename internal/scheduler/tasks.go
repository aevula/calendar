package scheduler

import (
	"time"

	"github.com/aevula/interview-hustlers-calendar/internal/config"
)

func PeriodicTasks(config.Config) []PeriodicTask {
	return []PeriodicTask{
		{
			Interval: 1 * time.Hour,
			Name:     "delete_old_events",
			Queue:    "delete_old_events_queue",
			Payload:  map[string]any{},
		},
	}
}
