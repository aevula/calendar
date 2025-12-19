package events

import "time"

type StartAtFilter struct {
	GTE time.Time
	LTE time.Time
}
