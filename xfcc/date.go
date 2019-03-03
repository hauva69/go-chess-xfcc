package xfcc

import (
	"log"
	"time"
)

// Date models the date.
type Date struct {
	Year  *int
	Month *int
	Day   *int
}

// NewDate returns a pointer to a new Date
func NewDate(year, month, day *int) *Date {
	return &Date{year, month, day}
}

// GetTime returns the Date as time.Time in some semi-sensible way, if applicable.
func (d *Date) GetTime() (time.Time, error) {
	layout := "2006.01.02"
	log.Fatalf("implement me: %s", layout)

	// FIXME
	return time.Now(), nil
}
