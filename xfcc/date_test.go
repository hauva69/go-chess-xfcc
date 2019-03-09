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
