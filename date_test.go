package xfcc

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const paulKeresWasBornInPGN = "1916.01.07"
const paulKeresWasBornInISO8601 = "1916-01-07"

// at least the ICCF implementation has no leading zeros
const paulKeresWasBornInXFCC = "1916.1.7"

const tigranPetrosianWasBornInPGN = "1929.06.17"
const tigranPetrosianWasBornInISO8601 = "1929-06-17"
const tigranPetrosianWasBornInXFCC = "1929.6.17"

const joséRaoulCapablancaWasBornInPgn = "1888.11.19"
const joséRaoulCapablancaWasBornInISO8601 = "1888-11-19"

func TestDatePGN(t *testing.T) {
	expected := "1916.01.07"
	year := 1916
	month := 1
	mday := 7
	d := NewDate(&year, &month, &mday)
	assert.Equal(t, expected, d.PGN())
}

func TestParsePGNDate(t *testing.T) {
	expectedYear := 1916
	expectedMonth := 1
	expectedDay := 7

	d, _ := Parse(paulKeresWasBornInPGN)
	assert.Equal(t, expectedYear, *d.Year)
	assert.Equal(t, expectedMonth, *d.Month)
	assert.Equal(t, expectedDay, *d.Day)
}

func TestParseUnParsableDate(t *testing.T) {
	_, err := Parse("200-3-4")
	if err == nil {
		t.Fatal("parsing 200-3-4 must fail")
	}
}

func TestParseNoMonthLeadingZero(t *testing.T) {
	expected := tigranPetrosianWasBornInPGN
	date, err := Parse(tigranPetrosianWasBornInXFCC)
	if err != nil {
		t.Fatalf("cannot parse %q: %s", tigranPetrosianWasBornInXFCC, err)
	}

	assert.Equal(t, expected, date.PGN())
}

func TestDatePGNYearIsNil(t *testing.T) {
	expected := "????.01.07"
	month := 1
	mday := 7
	d := NewDate(nil, &month, &mday)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNMonthIsNil(t *testing.T) {
	expected := "1916.??.07"
	year := 1916
	mday := 7
	d := NewDate(&year, nil, &mday)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNDayIsNil(t *testing.T) {
	expected := "1916.01.??"
	year := 1916
	month := 1
	d := NewDate(&year, &month, nil)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNDateIsNil(t *testing.T) {
	expected := "????.??.??"
	d := NewDate(nil, nil, nil)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNMonthAndDayAreNil(t *testing.T) {
	expected := "1916.??.??"
	year := 1916
	d := NewDate(&year, nil, nil)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNYearAndMonthAreNil(t *testing.T) {
	mday := 7
	expected := "????.??.07"
	d := NewDate(nil, nil, &mday)
	assert.Equal(t, expected, d.PGN())
}

func TestDatePGNYearAndDayAreNil(t *testing.T) {
	month := 1
	expected := "????.01.??"
	d := NewDate(nil, &month, nil)
	assert.Equal(t, expected, d.PGN())
}

func TestParseFromPGN(t *testing.T) {
	date, err := Parse(paulKeresWasBornInPGN)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, paulKeresWasBornInPGN, date.PGN())
}

func TestParseFromPGNYearIsNil(t *testing.T) {
	expected := "????.01.10"
	date, err := Parse(expected)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, date.PGN())
}

func TestParseFromPGNMonthIsNil(t *testing.T) {
	expected := "1916.??.07"
	date, err := Parse(expected)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, date.PGN())
}

func TestParseFromPGNDayIsNil(t *testing.T) {
	expected := "1916.01.??"
	date, err := Parse(expected)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, date.PGN())
}

func TestParseFromISO8601(t *testing.T) {
	date, err := Parse(paulKeresWasBornInISO8601)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, paulKeresWasBornInPGN, date.PGN())
}

func TestTime(t *testing.T) {
	paulKeresWasBornIn := time.Date(1916, time.January, 7, 0, 0, 0, 0, time.UTC)
	date, err := Parse(paulKeresWasBornInPGN)
	if err != nil {
		log.Fatal(err)
	}

	tmp, err := date.Time()
	if err != nil {
		log.Fatal(err)
	}

	assert.True(t, paulKeresWasBornIn.Equal(tmp))
}

func TestTimeYearIsUnknown(t *testing.T) {
	expected := "the year of PGN date \"????.01.07\" is nil"
	date, _ := Parse("????.01.07")
	_, err := date.Time()
	assert.Equal(t, expected, err.Error())
}

func TestTimeMonthIsUnknown(t *testing.T) {
	expected := "the month of PGN date \"1916.??.07\" is nil"
	date, _ := Parse("1916.??.07")
	_, err := date.Time()
	assert.Equal(t, expected, err.Error())
}

func TestTimeDayIsUnknown(t *testing.T) {
	expected := "the day of PGN date \"1916.01.??\" is nil"
	date, _ := Parse("1916.01.??")
	_, err := date.Time()
	assert.Equal(t, expected, err.Error())
}
