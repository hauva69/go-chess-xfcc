package xfcc

// ResultWhiteWins is the string value for a White win.
const ResultWhiteWins = "1-0"

// ResultBlackWins is the string value for a Black win.
const ResultBlackWins = "0-1"

// ResultDraw is the string value for a draw.
const ResultDraw = "1/2-1/2"

// ResultUnknown is the string value for an unknown result.
const ResultUnknown = "*"

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
func (r *Result) PGN() string {
	result := string(*r)
	switch result {
	case "WhiteWins":
		return ResultWhiteWins
	case "BlackDefaulted":
		return ResultWhiteWins
	case "WhiteWindAbjudicated":
		return ResultWhiteWins
	case "BlackWins":
		return ResultBlackWins
	case "WhiteDefaulted":
		return ResultBlackWins
	case "BlackWinAbjudicated":
		return ResultBlackWins
	case "Draw":
		return ResultDraw
	case "DrawAbjudicated":
		return ResultDraw
	case "WhiteWinAbjudicated":
		return ResultWhiteWins
	default:
		return ResultUnknown
	}
}
