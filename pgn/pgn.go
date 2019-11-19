package pgn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	xfcc "github.com/hauva69/go-chess-xfcc"
	"github.com/notnil/chess"
)

// Game returns the game from a PGN string
func Game(pgn string) (*chess.Game, error) {
	gameFunc, err := chess.PGN(strings.NewReader(pgn))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return chess.NewGame(gameFunc), nil
}

// FEN returns the final position of the game as a FEN string.
func FEN(pgn string) (string, error) {
	game, err := Game(pgn)
	if err != nil {
		return "", err
	}

	return game.FEN(), nil
}

// PieceCount returns the number of pieces on the board.
func PieceCount(pgn string) (int, error) {
	game, err := Game(pgn)
	if err != nil {
		return -1, err
	}

	return len(game.Position().Board().SquareMap()), nil
}

// EndGameResult returns true if the position is won in endgame tablebases.
// FIXME no fatals in a library!
// FIXME move this to a separate library
func EndGameResult(game xfcc.Game) bool {
	pgn, err := game.PGN()
	if err != nil {
		log.Fatal(err)
	}

	/*
		count, err := PieceCount(pgn)
		if err != nil {
			log.Fatal(err)
		}
		// FIXME replace the magic number

			if count > 7 {
				return false
			}
	*/

	tmpGame, err := Game(pgn)
	if err != nil {
		log.Fatal(err)
	}

	fen, err := FEN(pgn)
	if err != nil {
		log.Fatal(err)
	}

	fen = strings.ReplaceAll(fen, " ", "_")
	url := fmt.Sprintf("http://tablebase.lichess.ovh/standard/mainline?fen=%s", fen)
	log.Printf("URL=%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Print("too many requests, sleeping 1 minute")
		time.Sleep(time.Minute)
	} else if resp.StatusCode == http.StatusNotFound {
		log.Printf("404 Not found: %s", string(body))
		return false
	}

	log.Printf("BODY=%s", string(body))
	tmpGame.Position().Board().Draw()

	// FIXME
	return false
}
