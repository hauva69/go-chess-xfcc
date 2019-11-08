package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/docopt/docopt-go"
	"github.com/hauva69/go-chess-xfcc/configuration"
	"github.com/hauva69/go-chess-xfcc/xfcc"
)

func main() {
	usage := `Go-chess-xfcc.
	
	Usage:
		go-chess-xfcc
		go-chess-xfcc fetch
		go-chess-xfcc list [--myturn]
		go-chess-xfcc -h | --help
		go-chess-xfcc --version
  
  	Options:
		-h --help     Show this screen.
		--version     Show version
`

	arguments, err := docopt.Parse(usage, nil, true, "go-chess-xfcc 0.1", false)
	if err != nil {
		log.Fatalf("unable to parse arguments: %s", err)
	}

	config, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	if arguments["fetch"].(bool) {
		body, err := xfcc.GetMyGamesXML(xfcc.ICCFBaseURL, xfcc.ICCFSOAPMIMEType, config.User, config.Password)
		if err != nil {
			log.Fatalf("unable to get the XFCC XML: %s", err)
		}

		fmt.Print(string(body))
	} else if arguments["list"].(bool) {
		list(config, arguments["--myturn"].(bool))
	} else {
		games, err := xfcc.GetMyGames(xfcc.ICCFBaseURL, xfcc.ICCFSOAPMIMEType, config.User, config.Password)
		if err != nil {
			log.Fatal(err)
		}

		for _, game := range games {
			// TODO maybe --force option to bypass invalid data?
			pgn, err := game.PGN()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(pgn)
		}
	}
}

func list(config configuration.Configuration, myTurn bool) {
	games, err := xfcc.GetMyGames(xfcc.ICCFBaseURL, xfcc.ICCFSOAPMIMEType, config.User, config.Password)
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, game := range games {
		if !myTurn || (myTurn && game.MyTurn) {
			drawOffered := "–"
			if game.DrawOffered {
				drawOffered = "draw offered"
			}

			fmt.Fprintln(w, fmt.Sprintf(
				"%s\t%s\t%s\t%s",
				game.White,
				game.Black,
				game.Event,
				drawOffered,
			))
		}
	}

	w.Flush()
}
