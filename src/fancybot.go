package main

import (
	"context"
	"fmt"
	"log"

	"github.com/meganlee18/fancybot"
)

func main() {
	bot := slack.AuthenticateToSlack()
	slack.Respond(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Listening...")
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
