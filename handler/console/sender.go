package console

import (
	"errors"
	"isso0424/dise/handler/command"
	"log"
)

const messageTemplate = "channelID: %s message: %s\n"

type session struct{}

func (s session) Send(channelID string, message string) error {
	if len(message) > 2000 {
		return errors.New("too long message length")
	}

	if message == "" {
		return errors.New("cannot send empty message")
	}

	if channelID == "" {
		return errors.New("cannot send with empty channel ID")
	}
	log.Printf(messageTemplate, channelID, message)

	return command.ExecuteCommand(message, channelID, s)
}

func New() session {
	return session{}
}
