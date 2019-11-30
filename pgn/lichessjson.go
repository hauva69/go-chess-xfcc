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

	"github.com/notnil/chess"

	xfcc "github.com/hauva69/go-chess-xfcc"
)

type lichessResult struct {
	Mainline []move `json:"mainline"`
	Winner   string `json:"winner"`
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

// EndGameResult returns Win, Curse.
// FIXME move this to a separate library
// FIXME implement using both standard and standard/mainline endpoints
// (2) win, (1) cursed win, (0) draw, (-1) blessed loss, (-2) loss, (null) unknown
// WhiteWin, BlackWin, Unknown
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

	tmpGame, err := Game(pgn)
	if err != nil {
		return Unknown, err
	}

	fen, err := FEN(pgn)
	if err != nil {
		return Unknown, err
	}

	return EndGameResultFromFEN(tmpGame, fen)
}

// EndGameResultFromFEN returns the result of the position as a constant and 
// an error. Possible return values are WhiteWin, BlackWin and Unknown.
func EndGameResultFromFEN(game *chess.Game, fen string) (int, error) {
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

	game.Position().Board().Draw()
	var result lichessResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Unknown, err
	}

	if result.Winner == "w" {
		return WhiteWin, nil
	} else if result.Winner == "b" {
		return BlackWin, nil
	}

	// FIXME result.Winner when game is drawn

	return Unknown, nil
}
