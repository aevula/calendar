package queues

import "github.com/aevula/interview-hustlers-calendar/internal/config"

func Queues(config.Config) []Queue {
	return []Queue{
		NewQueue("delete_old_events_queue"),
		NewQueue("send_events_notifications_queue"),
	}
}
