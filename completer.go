package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

func completer(input prompt.Document) []prompt.Suggest {
	if input.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}

	args := strings.Split(input.TextBeforeCursor(), " ")

	suggestions := []prompt.Suggest{
		{Text: "full-deploy", Description: "pull, full shutdown, and restart of environment"},
		{Text: "half-deploy", Description: "pull and quick restart of environment"},
	}

	// Only show suggestions if its the first command
	if len(args) > 1 {
		return []prompt.Suggest{}
	}

	return prompt.FilterHasPrefix(suggestions, input.GetWordBeforeCursor(), true)
}
