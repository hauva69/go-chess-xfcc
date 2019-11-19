package pgn

import (
	"log"
	"strings"

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
func EndGameResult(game xfcc.Game) bool {
	pgn, err := game.PGN()
	if err != nil {
		log.Fatal(err)
	}

	count, err := PieceCount(pgn)
	if err != nil {
		log.Fatal(err)
	}
	// FIXME replace the magic number
	if count > 7 {
		return false
	}

	tmpGame, err := Game(pgn)
	if err != nil {
		log.Fatal(err)
	}

	tmpGame.Position().Board().Draw()

	// FIXME
	return false
}
