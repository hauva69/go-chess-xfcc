package xfcc

import (
	"fmt"
	"log"
	"time"
)

const summerSolstice1916 = "1916-06-21"

// PGNLayout is date layout for PGN tags.
const PGNLayout = "2006.01.02"

// Date models the date.
type Date struct {
	Year  *int
	Month *int
	Day   *int
}

// NewDate returns a pointer to a new Date.
func NewDate(year, month, day *int) *Date {
	return &Date{year, month, day}
}

// GetTime returns the Date as time.Time in some semi-sensible way, if applicable.
// Summer solstice of certain, fixed year might be a reasonable option.
func (d *Date) GetTime() (time.Time, error) {
	layout := "2006.01.02"
	log.Fatalf("implement me: %s", layout)

	// FIXME
	return time.Now(), nil
}

// PGN returns the date in PGN format.
func (d *Date) PGN() string {
	s := "????."

	if d.Year != nil {
		s = fmt.Sprintf("%04d.", *d.Year)
	}

	if d.Month == nil {
		s = fmt.Sprintf("%s??.", s)
	} else {
		s = fmt.Sprintf("%s%02d.", s, *d.Month)
	}

	if d.Day == nil {
		s = fmt.Sprintf("%s??", s)
	} else {
		s = fmt.Sprintf("%s%02d", s, *d.Day)
	}

	return s
}
