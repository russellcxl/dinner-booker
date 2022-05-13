package main

import (
	"flag"
	"git.garena.com/russell.chanxl/dinner/cli"
	"git.garena.com/russell.chanxl/dinner/http_"
	"git.garena.com/russell.chanxl/dinner/utils"
	"log"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "token", "", "input SEA dinner token")
	flag.Parse()
}

func main() {
	var t string
	if Token == "" {
		log.Printf("did not input token; attempting to retrive from .env file\n")
		var err error
		t, err = utils.GetTokenFromLocalEnv()
		if err != nil {
			log.Fatalf("failed to retrieve token from .env file: %s\n", err.Error())
		}
	} else {
		t = Token
	}

	bot, err := http_.NewDinnerBot("https://dinner.sea.com/", t)
	if err != nil {
		log.Fatalf("failed to initialise dinner bot: %s\n", err.Error())
	}

	cli.Execute(bot)
}