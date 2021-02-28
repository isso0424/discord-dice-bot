package console

import (
	"errors"
	"log"
)

type session struct {}

func (s session) Send(channelID string, message string) error {
	if len(message) > 2000 {
		return errors.New("too long message length")
	}

	if message == "" {
		return errors.New("cannot send empty message")
	}
	log.Println(message)

	return nil
}
