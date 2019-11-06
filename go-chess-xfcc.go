package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/hauva69/go-chess-xfcc/configuration"
	"github.com/hauva69/go-chess-xfcc/xfcc"
)

func main() {
	usage := `Go-chess-xfcc.
	
	go-chess-xfcc
	`

	arguments, err := docopt.Parse(usage, nil, true, "go-chess-xfcc 0.1", false)
	log.Printf("arguments=%+v", arguments)

	config, err := configuration.GetConfiguration()
	if err != nil {
		log.Fatal(err)
	}

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
