package slack

import "github.com/shomali11/slacker"

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
