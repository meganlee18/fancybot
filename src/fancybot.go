package main

import (
	"context"
	"log"

	"github.com/meganlee18/fancybot"
)

func main() {
	bot := slack.AuthenticateToSlack()
	slack.Respond(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Listening...")

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ready! Say something to FancyBot on Slack!")
}
