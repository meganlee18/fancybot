package slack

import (
	"github.com/shomali11/slacker"
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

func threadReplyHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.ReportError(errors.New("oops"), slacker.WithThreadError(true))
}

func messageReplyHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.ReportError(errors.New("oops"))
}

func repeatWordHandle(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	word := request.StringParam("word", "Hello!")
	number := request.IntegerParam("number", 1)

	for i := 0; i < number; i++ {
		response.Reply(word)
	}
}
