package xfcc

import "testing"

const paulKeresWasBornInPGN = "1916.01.07"
const paulKeresWasBornInISO8601 = "1916-01-07"

func TestEventDatePgn(t *testing.T) {
	ed := NewEventDate()
	t.Fatalf("FIXME: %+v", ed)
}
