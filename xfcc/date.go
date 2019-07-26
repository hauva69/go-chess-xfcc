package xfcc

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

const summerSolstice1916 = "1916-06-21"

// DateRegularExpressionString is the string that is used to
// create DateRegularExpression.
const DateRegularExpressionString = `^(\d\d\d\d)[.-](\d\d?)[.-](\d\d?)$`

// DateRegularExpression is used for parsin Dates from strings.
var DateRegularExpression *regexp.Regexp

func init() {
	DateRegularExpression = regexp.MustCompile(DateRegularExpressionString)
}

// PGNLayout is date layout for PGN tags.
const PGNLayout = "2006.01.02"

// ISO8601Layout is date layout for ISO8601 standard.
const ISO8601Layout = "2006-01-02"

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

// Now returns a pointer to the current day.
func Now() *Date {
	tm := time.Now()
	year := tm.Year()
	month := int(tm.Month())
	day := int(tm.Day())

	return NewDate(&year, &month, &day)
}

// Parse returns a pointer to a new Date from string which is formatted like
// "1916.01.07", "1916-01-07", 1916-1-7 or "1916.1.7" and an error.
func Parse(s string) (*Date, error) {
	groups := DateRegularExpression.FindAllStringSubmatch(s, -1)
	log.Printf("parsing %q, got %+v", s, groups)
	if len(groups) != 1 || len(groups[0]) != 4 {
		return nil, fmt.Errorf(
			"parsing the date failed: %q does not match %q",
			s,
			DateRegularExpressionString,
		)
	}

	year, err := strconv.Atoi(groups[0][1])
	if err != nil {
		return nil, err
	}

	month, err := strconv.Atoi(groups[0][2])
	if err != nil {
		return nil, err
	}

	day, err := strconv.Atoi(groups[0][3])
	if err != nil {
		return nil, err
	}

	return NewDate(&year, &month, &day), nil
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
