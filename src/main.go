package main

import (
	"../src/fancybot"
)

func main() {
	rtm := fancybot.SlackConnect()
	fancybot.SendAndReceiveEvents(rtm)
}
