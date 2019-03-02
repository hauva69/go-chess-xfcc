package xfcc

import "testing"

const PaulKeresWasBornInPGN = "2016.01.07"
const PaulKeresWasBornInISO8601 = "2016-01-07"

func TestEventDatePgn(t *testing.T) {
	ed := NewEventDate()
	t.Fatalf("FIXME: %+v", ed)
}
