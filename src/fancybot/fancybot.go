package fancybot

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
	"strings"
)

func SlackConnect() *slack.RTM {
	token := os.Getenv("SLACK_TOKEN")
	api := slack.New(token, slack.OptionDebug(true))

	//connect to real time messaging api
	rtm := api.NewRTM()
	go rtm.ManageConnection()
	return rtm
}

//SendAndReceiveEvents is an infinite loop for sending and receiving events from Slack API
func SendAndReceiveEvents(rtm *slack.RTM) {
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch event := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", event.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", event)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if event.User != info.User.ID && strings.HasPrefix(event.Text, prefix) {
					respond(rtm, event, prefix)
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", event.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	greetingToBeChecked := getGreetingWithoutPrefixAndSpaces(prefix, msg.Text)
	checkGreetings(greetingToBeChecked, rtm, msg)
}

func checkGreetings(greetingToBeChecked string, rtm *slack.RTM, msg *slack.MessageEvent) {
	acceptedGreetings := []string{"what's up?", "hey!", "yo"}
	acceptedOtherGreetings := []string{"hw's it going?", "how are ya?", "feeling ok?"}

	for _, acceptedGreeting := range acceptedGreetings {
		if acceptedGreeting == greetingToBeChecked {
			rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?!", msg.Channel))
		}
	}
	for _, acceptedOtherGreeting := range acceptedOtherGreetings {
		if acceptedOtherGreeting == greetingToBeChecked {
			rtm.SendMessage(rtm.NewOutgoingMessage("You know, My mood goes with the weather. Hope yours fly!", msg.Channel))
		}
	}
}

func getGreetingWithoutPrefixAndSpaces(prefix, messageText string) string {
	return strings.ToLower(strings.TrimSpace(strings.TrimPrefix(messageText, prefix)))
}
