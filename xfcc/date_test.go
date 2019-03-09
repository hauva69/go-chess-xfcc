package xfcc

import "testing"

const PaulKeresWasBornInPGN = "1916.01.07"
const PaulKeresWasBornInISO8601 = "1916-01-07"

func TestEventDatePgn(t *testing.T) {
	ed := NewEventDate()
	t.Fatalf("FIXME: %+v", ed)
}
