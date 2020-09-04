package slack

import (
	"os"

	"github.com/shomali11/slacker"
)

func AuthenticateToSlack() *slacker.Slacker {
	token := os.Getenv("SLACK_TOKEN")
	return slacker.NewClient(token)
}

func Respond(bot *slacker.Slacker) {
	bot.Command("Hello <name>", helloDefinition())
	bot.Command("ping", pingDefinition())
	bot.Command("repeat <word> <number>", repeatWordDefinition())
	bot.Command("thread", threadReplyDefinition())
	bot.Command("message", messageReplyDefinition())
}
