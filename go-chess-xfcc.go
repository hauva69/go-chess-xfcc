package main

import (
	"fmt"
	"log"

	"github.com/hauva69/go-chess-xfcc/configuration"
	"github.com/hauva69/go-chess-xfcc/xfcc"
)

func main() {
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
