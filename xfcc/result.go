package xfcc

// Result models the result of the game. It can have following values:
// - "Ongoing"
// - "WhiteWins"
// - "BlackWins"
// - "Draw"
// - "WhiteWinAdjudicated"
// - "BlackWinAdjudicated"
// - "DrawAdjudicated"
// - "WhiteDefaulted"
// - "BlackDefaulted"
// - "BothDefaulted"
// - "Cancelled"
// - "AdjudicationPending"
type Result string

// PGN returns the result of the game as a PGN string.
func (r *Result) PGN() (string, error) {
	/*
		"Ongoing"
		"WhiteWins"
		"BlackWins"
		"Draw"
		"WhiteWinAdjudicated"
		"BlackWinAdjudicated"
		"DrawAdjudicated"
		"WhiteDefaulted"
		"BlackDefaulted"
		"BothDefaulted"
		"Cancelled"
		"AdjudicationPending"
	*/

	return string(*r), nil
}
