package fancybot

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type greetingsTest struct {
	prefix string
	message string
	expected string
}

var greets = []greetingsTest{
	{
		prefix:  "<@U1A2B3C4D>",
		message: "hey!",
		expected: "hey!",
	},
	{
		prefix:  "<@W1A2B3C4D>",
		message: "how's it going?",
		expected: "how's it going?",
	},
	{
		prefix:  "<@345435345>",
		message: "feeling ok?",
		expected: "feeling ok?",
	},
}

func Test_getGreetingWithoutPrefixAndSpaces(t *testing.T) {
	for _, greet := range greets {
		actual := getGreetingWithoutPrefixAndSpaces(greet.prefix, greet.message)
		assert.Equal(t, greet.expected, actual)
	}
}
