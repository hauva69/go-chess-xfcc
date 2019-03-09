package xfcc

import "testing"

const paulKeresWasBornInPGN = "1916.01.07"
const paulKeresWasBornInISO8601 = "1916-01-07"

// at least the ICCF implementation has no leading zeros
const paulKeresWasBornInXFCC = "1916-1-7"

func TestEventDatePgn(t *testing.T) {
	ed := NewEventDate()
	t.Fatalf("FIXME: %+v", ed)
}
