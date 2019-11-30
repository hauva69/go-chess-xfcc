package pgn

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	xfcc "github.com/hauva69/go-chess-xfcc"
)

type lichessResult struct {
	Mainline []move `json:"mainline"`
	Winner   *string `json:"winner"`
}

type move struct {
	UCI string `json:"uci"`
	DTZ int    `json:"dtz"`
}

// MaximumPiecesForEndgameResolving is the maximum number of pieces 
// of endgame tablebases.
const MaximumPiecesForEndgameResolving = 7

// FIXME make this configurable
const mainlineURL = "http://tablebase.lichess.ovh/standard/mainline?fen"

const (
	// Loss is a loss according to the Syzygy table bases.
	Loss = (-2 + iota)
	// BlessedLoss is a blessed loss according to the Syzygy table bases.
	BlessedLoss
	// Draw is a draw according to the Syzygy table bases.
	Draw
	// CursedWin is a cursed win according to the Syzygy table bases.
	CursedWin
	// Win is a win according to the Syzygy table bases.
	Win
	// Unknown is for an unknown result.
	Unknown
	// WhiteWin White wins.
	WhiteWin
	// BlackWin Black wins.
	BlackWin
)

// EndGameResult returns WhiteWin, BlackWin, Draw or Unknown and an error.
// FIXME move this to a separate library
// FIXME implement using both standard and standard/mainline endpoints
// (2) win, (1) cursed win, (0) draw, (-1) blessed loss, (-2) loss, (null) unknown
func EndGameResult(game xfcc.Game) (int, error) {
	pgn, err := game.PGN()
	if err != nil {
		return Unknown, err
	}

	count, err := PieceCount(pgn)
	if err != nil {
		return Unknown, err
	}

	if count > MaximumPiecesForEndgameResolving {
		msg := fmt.Sprintf("number of pieces is %d, no table bases exist", count)
		return Unknown, fmt.Errorf(msg)
	}

	fen, err := FEN(pgn)
	if err != nil {
		return Unknown, err
	}

	return EndGameResultFromFEN(fen)
}

// EndGameResultFromFEN returns the result of the position as a constant and 
// an error. Possible return values are WhiteWin, BlackWin, Draw and Unknown.
func EndGameResultFromFEN(fen string) (int, error) {
	fen = strings.ReplaceAll(fen, " ", "_")
	url := fmt.Sprintf("%s=%s", mainlineURL, fen)
	resp, err := http.Get(url)
	if err != nil {
		return Unknown, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Unknown, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Print("too many requests, sleeping 1 minute")
		time.Sleep(time.Minute)
	} else if resp.StatusCode == http.StatusNotFound {
		msg := fmt.Sprintf("404 Not found: %s", string(body))
		return Unknown, errors.New(msg)
	}

	var result lichessResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Unknown, err
	}

	if nil == result.Winner {
		return Draw, nil
	} else if "w" == *result.Winner {
		return WhiteWin, nil
	} else if "b" == *result.Winner {
		return BlackWin, nil
	} 

	return Unknown, nil
}
