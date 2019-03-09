package xfcc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const paulKeresWasBornInPGN = "1916.01.07"
const paulKeresWasBornInISO8601 = "1916-01-07"

// at least the ICCF implementation has no leading zeros
const paulKeresWasBornInXFCC = "1916-1-7"

func TestDatePGN(t *testing.T) {
	expected := "1916.01.07"
	year := 1916
	month := 1
	mday := 7
	d := NewDate(&year, &month, &mday)
	assert.Equal(t, expected, d.PGN())
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
