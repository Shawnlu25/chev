package event

import "time"

// Event is a time log event
type Event struct {
	StartTime   time.Time
	Duration    time.Duration
	Description string
}
