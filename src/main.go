package main

import (
	"github.com/meganlee18/fancybot/src/fancybot"
)

func main() {
	rtm := fancybot.SlackConnect()
	fancybot.SendAndReceiveEvents(rtm)
}
