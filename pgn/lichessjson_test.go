package pgn

import (
	"testing"

    "github.com/stretchr/testify/assert"
    
    xfcc "github.com/hauva69/go-chess-xfcc"
)

func TestWin(t *testing.T) {
	const expected = 2
	assert.Equal(t, expected, Win)
}

func TestCursedWin(t *testing.T) {
	const expected = 1
	assert.Equal(t, expected, CursedWin)
}

func TestDraw(t *testing.T) {
	const expected = 0
	assert.Equal(t, expected, Draw)
}

func TestBlessedLoss(t *testing.T) {
	const expected = -1
	assert.Equal(t, expected, BlessedLoss)
}

func TestLoss(t *testing.T) {
	const expected = -2
	assert.Equal(t, expected, Loss)
}

func TestWhiteWins(t *testing.T) {
    game := xfcc.Game{
        Movetext: WhiteWinsMovetext, 
        Date: "2019.04.10", 
        EventDate: "2019.04.10",
    }

    result, _ := EndGameResult(game)
    assert.Equal(t, WhiteWin, result, "Winner must be WhiteWin.")
}

func TestBlackWins(t *testing.T) {
    game := xfcc.Game{
        Movetext: BlackWinsMovetext, 
        Date: "2018.10.01", 
        EventDate: "2018.10.01",
    }

    result, _ := EndGameResult(game)
    assert.Equal(t, BlackWin, result, "Winner must be BlackWin.")
}

// bareKingsDrawFEN is a FEN for position where there are only the two kings.
const bareKingsDrawFEN = "4k3/8/8/8/8/8/8/4K3 w - - 0 1"

func TestBareKingsDraw(t *testing.T) {
    result, err := EndGameResultFromFEN(bareKingsDrawFEN)
    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, Draw, result, "The result must be a draw")
}

func TestStartingPositionEndGameResult(t *testing.T) {
    result, err := EndGameResultFromFEN(StartingPositionFEN)
    if err != nil {
        t.Log(err)
    }

    assert.Equal(t, Unknown, result, "The result must be Unknown.")
}

// FewPiecesGame1 is something to be fixed
// FIXME remove links to the test data
// FIXME move these to a separate test file
// https://www.iccf.com/game?id=1086957[Event "FIN/AC/4 (FIN)"[Date "2019.04.10"]
// [Round "-"]
// [White "Mäkelä, Ari"]
// [Black "Rajala, Olavi"]
// [Result "1-0"]
// [WhiteElo "2181"]
// [BlackElo "1995"]
const WhiteWinsMovetext = `1.Nf3 e6 2.c4 Nf6 3.Nc3 c5 4.g3 Nc6 5.Bg2 Qb6 6.O-O Be7 7.b3 d5 8.Na4 Qa6 9.Bb2 O-O 10.Bxf6 Bxf6 11.cxd5 exd5 12.Nxc5 Qa5 13.Rc1 Bb2 14.Rb1 Ba3 15.d4 Bf5 16.Nd3 Rac8 17.Nh4 Bxd3 18.Qxd3 Nb4 19.Qf3 g6 20.Bh3 Rc2 21.Ng2 Rd8 22.Ne3 Rc6 23.Bg2 Qc7 24.h4 b5 25.h5 f5 26.Nc4 dxc4 27.bxc4 Rxc4 28.hxg6 hxg6 29.Qxa3 a5 30.Qf3 Nc2 31.d5 Kg7 32.Qd3 Qc5 33.Rfc1 Nd4 34.Rxc4 Qxc4 35.Qxc4 bxc4 36.Kf1 Kf6 37.Rb6+ Kf7 38.Rb7+ Kf6 39.e3 Nc2 40.a4 c3 41.Rc7 Rb8 42.Rxc3 Rb1+ 43.Ke2 Nb4 44.f4 Rb2+ 45.Kf1 Rb1+ 46.Kf2 Rb2+ 47.Kg1 Ra2 48.e4 fxe4 49.Bxe4 Rxa4 50.Rc4 Ra2 51.Rd4 Nc2 52.Rd1 Ke7 53.d6+ Kd8 54.Bc6 Nb4 55.Bb5 Rc2 56.Re1 Nc6 57.Re6 Rc5 58.Bf1 a4 59.Rxg6 Rc1 60.Kf2 Kc8 61.Bh3+ Kb7 62.Rg7+ Kb6 63.Rg8 a3 64.Ra8 Rc2+ 65.Ke3 Rc3+ 66.Ke4 Rxg3 67.Be6 Rg6 68.Kf5 Rg3 69.d7 Rd3 70.Ke4 Nb4 71.Rb8+ Kc5 72.Rxb4 Rxd7 73.Bxd7 Kxb4 1-0`

// FewPiecesGame2 is something to be fixed
// https://www.iccf.com/game?id=1049523
// [Event "FIN/WS/47 (FIN)"]
// [Site "ICCF"]
// [Date "2018.10.01"]
// [Round "-"]
// [White "Hoogkamer, Michael"]
// [Black "Mäkelä, Ari"]
// [Result "0-1"]
// [WhiteElo "1975"]
const BlackWinsMovetext = `
1.d4 d5 2.c4 c6 3.cxd5 cxd5 4.Bf4 Nc6 5.e3 Qb6 6.Nc3 Nf6 7.Bd3 a6 8.Nge2 e6 9.O-O Be7 10.Na4 Qd8 11.h3 Nb4 12.Bb1 Nc6 13.Bd3 Nb4 14.Bb1 Nc6 15.a3 Na5 16.Qc2 Nc4 17.b3 Nd6 18.Rc1 Bd7 19.Nc5 Bc6 20.Qb2 O-O 21.Nd3 Qb6 22.a4 a5 23.g4 Rfc8 24.f3 Be8 25.Nc5 e5 26.dxe5 Nc4 27.Qc3 Nxe5 28.Bxe5 Bxc5 29.Bd4 Bxd4 30.Qxd4 Qxb3 31.Rxc8 Rxc8 32.g5 Nd7 33.Ba2 Qc2 34.Nf4 Qc7 35.Kg2 Kf8 36.Bxd5 Nb6 37.Be4 Rd8 38.Rc1 Qxc1 39.Qxd8 Nc4 40.Nd5 Qb2+ 41.Kg3 Qe5+ 42.Kg2 Qd6 43.Qxd6+ Nxd6 44.Bc2 b5 45.Nb6 bxa4 46.Nxa4 Nc4 47.Kf2 Bd7 48.h4 Bxa4 49.Bxa4 Nb6 50.Bb3 a4 51.Ba2 Nc8 52.Ke1 Nd6 53.Bb1 a3 54.Kd2 Nf5 55.h5 h6 56.f4 Ng3 57.Kc3 Nxh5 58.Bf5 Ng3 59.gxh6 gxh6 60.Bh3 Ne4+ 61.Kb3 h5 62.Kxa3 Nf2 63.Bf1 h4 64.Kb3 h3 65.Bxh3 Nxh3 66.Kc2 Ke7 0-1
`

const WhiteWinsJSON = `
{
    "mainline": [
        {
            "uci": "d7e6",
            "dtz": -2
        },
        {
            "uci": "b4c3",
            "dtz": 1
        },
        {
            "uci": "f4f5",
            "dtz": -2
        },
        {
            "uci": "c3b2",
            "dtz": 1
        },
        {
            "uci": "f5f6",
            "dtz": -2
        },
        {
            "uci": "b2a1",
            "dtz": 1
        },
        {
            "uci": "f6f7",
            "dtz": -2
        },
        {
            "uci": "a1b1",
            "dtz": 1
        },
        {
            "uci": "f7f8q",
            "dtz": -4
        },
        {
            "uci": "b1b2",
            "dtz": 3
        },
        {
            "uci": "e6a2",
            "dtz": -2
        },
        {
            "uci": "b2a1",
            "dtz": 1
        },
        {
            "uci": "f8b4",
            "dtz": -1
        },
        {
            "uci": "a1a2",
            "dtz": 3
        },
        {
            "uci": "b4c3",
            "dtz": -2
        },
        {
            "uci": "a2b1",
            "dtz": 1
        },
        {
            "uci": "c3a3",
            "dtz": -8
        },
        {
            "uci": "b1c2",
            "dtz": 7
        },
        {
            "uci": "a3b4",
            "dtz": -6
        },
        {
            "uci": "c2d1",
            "dtz": 5
        },
        {
            "uci": "b4b2",
            "dtz": -4
        },
        {
            "uci": "d1e1",
            "dtz": 3
        },
        {
            "uci": "e4e3",
            "dtz": -2
        },
        {
            "uci": "e1d1",
            "dtz": 1
        },
        {
            "uci": "b2b1",
            "dtz": -1
        }
    ],
    "winner": "w",
    "dtz": 3
}`

const BlackWinsJSON = `{
    "mainline": [
        {
            "uci": "c2d1",
            "dtz": 4
        },
        {
            "uci": "h3f2",
            "dtz": -3
        },
        {
            "uci": "d1e1",
            "dtz": 2
        },
        {
            "uci": "f2e4",
            "dtz": -1
        },
        {
            "uci": "e1d1",
            "dtz": 1
        },
        {
            "uci": "f7f6",
            "dtz": -1
        },
        {
            "uci": "d1c1",
            "dtz": 1
        },
        {
            "uci": "f6f5",
            "dtz": -11
        },
        {
            "uci": "c1b1",
            "dtz": 10
        },
        {
            "uci": "e4d6",
            "dtz": -9
        },
        {
            "uci": "b1c2",
            "dtz": 8
        },
        {
            "uci": "e7e6",
            "dtz": -7
        },
        {
            "uci": "c2d1",
            "dtz": 6
        },
        {
            "uci": "e6d5",
            "dtz": -5
        },
        {
            "uci": "d1e1",
            "dtz": 4
        },
        {
            "uci": "d6c4",
            "dtz": -3
        },
        {
            "uci": "e1e2",
            "dtz": 2
        },
        {
            "uci": "d5e4",
            "dtz": -1
        },
        {
            "uci": "e2d1",
            "dtz": 1
        },
        {
            "uci": "c4e3",
            "dtz": -2
        },
        {
            "uci": "d1c1",
            "dtz": 1
        },
        {
            "uci": "e4f4",
            "dtz": -4
        },
        {
            "uci": "c1b1",
            "dtz": 3
        },
        {
            "uci": "f4f3",
            "dtz": -2
        },
        {
            "uci": "b1a1",
            "dtz": 1
        },
        {
            "uci": "f5f4",
            "dtz": -4
        },
        {
            "uci": "a1b1",
            "dtz": 3
        },
        {
            "uci": "f3e2",
            "dtz": -2
        },
        {
            "uci": "b1a1",
            "dtz": 1
        },
        {
            "uci": "f4f3",
            "dtz": -2
        },
        {
            "uci": "a1b1",
            "dtz": 1
        },
        {
            "uci": "f3f2",
            "dtz": -2
        },
        {
            "uci": "b1a1",
            "dtz": 1
        },
        {
            "uci": "f2f1q",
            "dtz": -8
        },
        {
            "uci": "a1a2",
            "dtz": 7
        },
        {
            "uci": "e3c4",
            "dtz": -6
        },
        {
            "uci": "a2b3",
            "dtz": 5
        },
        {
            "uci": "f1a1",
            "dtz": -4
        },
        {
            "uci": "b3b4",
            "dtz": 3
        },
        {
            "uci": "a1a5",
            "dtz": -2
        },
        {
            "uci": "b4b3",
            "dtz": 1
        },
        {
            "uci": "e2d1",
            "dtz": -1
        },
        {
            "uci": "b3c4",
            "dtz": 11
        },
        {
            "uci": "a5e5",
            "dtz": -10
        },
        {
            "uci": "c4d3",
            "dtz": 9
        },
        {
            "uci": "d1c1",
            "dtz": -8
        },
        {
            "uci": "d3c4",
            "dtz": 7
        },
        {
            "uci": "c1c2",
            "dtz": -6
        },
        {
            "uci": "c4b4",
            "dtz": 5
        },
        {
            "uci": "e5d5",
            "dtz": -4
        },
        {
            "uci": "b4a4",
            "dtz": 3
        },
        {
            "uci": "c2c3",
            "dtz": -2
        },
        {
            "uci": "a4a3",
            "dtz": 1
        },
        {
            "uci": "d5b3",
            "dtz": -1
        }
    ],
    "winner": "b",
    "dtz": -5
}`
