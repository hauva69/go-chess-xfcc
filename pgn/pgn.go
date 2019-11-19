package pgn

import (
	"log"
	"strings"

	"github.com/notnil/chess"
)

// FEN returns the final position of the game as a FEN string.
func FEN(pgn string) (string, error) {
	gameFunc, err := chess.PGN(strings.NewReader(pgn))
	if err != nil {
		log.Print(err)
		return "ERROR", err
	}

	return chess.NewGame(gameFunc).FEN(), nil
}
