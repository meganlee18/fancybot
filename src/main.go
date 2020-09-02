package main

import (
	"context"
	"fmt"
	"github.com/shomali11/slacker"
	"log"
	"os"
	"errors"
)

func helloHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	name := request.Param("name")
	if name == "" {
		response.Reply("Hello stranger")
		return
	}

	response.Reply("Hey " + name + "!")
}

func pingHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.Reply("pong")

}

func repeatWordHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	word := request.StringParam("word", "Hello!")
	number := request.IntegerParam("number", 1)

	for i := 0; i < number; i++ {
		response.Reply(word)
	}

}

func helloDefinition() *slacker.CommandDefinition {
		return &slacker.CommandDefinition{
			Description: "Echo hello with name",
			Example: "Hello, or Hello Ben",
			Handler: helloHandle,
		}
}

func pingDefinition() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description : "Ping!",
		Example: "ping",
		Handler: pingHandle,
	}
}

func repeatWordDefinition() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description : "Repeat a word a number of times!",
		Example: "repeat tree 10",
		Handler: repeatWordHandle,
	}
}

func threadReplyHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.ReportError(errors.New("oops"), slacker.WithThreadError(true))
}

func messageReplyHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.ReportError(errors.New("oops"))
}

func threadReplyDefinition() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description : "Tests errors in new threads",
		Handler: threadReplyHandle,
	}
}

func messageReplyDefinition() *slacker.CommandDefinition {
	return &slacker.CommandDefinition{
		Description : "Tests errors in new messages",
		Handler: messageReplyHandle,
	}
}

func authenticateToSlack() *slacker.Slacker{
	token := os.Getenv("SLACK_TOKEN")
	return slacker.NewClient(token)
}

func main() {
	bot := authenticateToSlack()
	bot.Command("Hello <name>", helloDefinition())
	bot.Command("ping", pingDefinition())
	bot.Command("repeat <word> <number>", repeatWordDefinition())
	bot.Command("thread", threadReplyDefinition())
	bot.Command("message", messageReplyDefinition())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("Listening...")
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
