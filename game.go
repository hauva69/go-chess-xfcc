package xfcc

// FIXME Check the difference, or lack of it, of Date and EventDate when the next tournament begins.

import (
	"encoding/xml"
	"fmt"

	"github.com/eidolon/wordwrap"
)

// Player models a chess player.
type Player string

// MaximumMoveTextLength defines the maximum length of lines of the movetext of a game.
const MaximumMoveTextLength = 78

// Movetext models the movetext of the game.
type Movetext string

// Wrap returns the movetext as a wrapped string, see MaximumMoveTextLength.
func (mt *Movetext) Wrap() string {
	wrapper := wordwrap.Wrapper(MaximumMoveTextLength, false)

	return wrapper(string(*mt))
}

// Game models a correspondence chess game.
type Game struct {
	XMLName     xml.Name `xml:"XfccGame"`
	ID          uint64   `xml:"id"`
	Event       string   `xml:"event"`
	Site        string   `xml:"site"`
	Date        string   `xml:"date"`
	EventDate   string   `xml:"eventDate"`
	White       Player   `xml:"white"`
	Black       Player   `xml:"black"`
	Result      Result   `xml:"result"`
	Movetext    Movetext `xml:"moves"`
	WhiteElo    int      `xml:"whiteElo"`
	BlackElo    int      `xml:"blackElo"`
	DrawOffered bool     `xml:"drawOffered"`
	MyTurn      bool     `xml:"myTurn"`
}

// PGN returns the game as a PGN string.
func (g *Game) PGN() (string, error) {
	date, err := Parse(g.EventDate)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
			PGNTemplate,
			g.Event,
			g.Site,
			date.PGN(),
			g.White,
			g.Black,
			g.Result.PGN(),
			g.WhiteElo,
			g.BlackElo,
			g.Movetext.Wrap(),
			g.Result.PGN(),
		),
		nil
}

// PGNTemplate is a template for Portable Game Notation.
const PGNTemplate = `[Event "%s"]
[Site "%s"]
[Date "%s"]
[Round "-"]
[White "%s"]
[Black "%s"]
[Result "%s"]
[WhiteElo "%d"]
[BlackElo "%d"]

%s %s
`
