package parser

import (
	"errors"
	"strings"
)

func ParseCommand(text string) (command string, args []string, err error) {
	texts := strings.Split(text, " ")
	if len(texts) == 0 {
		err = errors.New("command not provides")

		return
	}

	command = texts[0]
	if len(texts) < 2 {

		return
	}

	args = texts[1:]

	return
}
