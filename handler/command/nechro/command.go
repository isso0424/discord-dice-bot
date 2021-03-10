package nechro

import (
	"errors"
	"fmt"
	"isso0424/dise/types"
	"strconv"
)

type Nechro struct{}

func (cmd Nechro) Exec(channelID string, args []string, session types.Session) error {
	if !cmd.ValidateArgs(args) {
		return errors.New(fmt.Sprintf("invalid args: %v", args))
	}

	amendment, _ := strconv.Atoi(args[0])

	message, _, _ := judge(amendment)

	return session.Send(channelID, message)
}

func (cmd Nechro) GetPrefix() string {
	return "!nechro"
}

func (cmd Nechro) GetHelp() string {
	return "修正値を基にネクロニカの行動判定をします。"
}

func (cmd Nechro) ValidateArgs(args []string) bool {
	if len(args) < 1 {
		return false
	}

	_, err := strconv.Atoi(args[0])
	if err != nil {
		return false
	}

	return true
}
