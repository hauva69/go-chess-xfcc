package xfcc

import (
	"log"
	"time"
)

// EventDate models the starting date of an event.
// FIXME rename me to PGNDate
type EventDate struct {
	time.Time
}

// NewEventDate returns a pointer to a new EventDate
// FIXME take year, month and date as parameters
func NewEventDate() *EventDate {
	return &EventDate{}
}

// GetTime returns the the as a time.Time in some semi-sensible way, if applicable.
func (d *EventDate) GetTime(time.Time) {
	layout := "2006.01.02"
	log.Fatalf("implement me: %s", layout)
}
