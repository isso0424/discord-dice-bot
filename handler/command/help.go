package command

import (
	"fmt"
	"isso0424/dise/types"
)

type help struct{}

func (cmd help) GetHelp() string {
	return "コマンド一覧を表示します"
}

func (cmd help) GetPrefix() string {
	return "!help"
}

func (cmd help) ValidateArgs(args []string) bool {
	return true
}

func (cmd help) Exec(channelID string, args []string, session types.Session) error {
	return session.Send(channelID, createHelpMessage())
}

func createHelpMessage() (message string) {
	for _, command := range commands {
		message += fmt.Sprintf("%s: %s\n", command.GetPrefix(), command.GetHelp())
	}

	return
}
