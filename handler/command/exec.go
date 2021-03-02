package command

import (
	"fmt"
	"isso0424/dise/types"
)

func Exec(cmd Command, channelID string, args []string, session types.Session) error {
	if !cmd.ValidateArgs(args) {
		session.Send(channelID, "引数が間違っています。")
		return fmt.Errorf("invalid arguments: %v target: %s", args, cmd.GetPrefix())
	}

	err := cmd.Exec(channelID, args, session)
	if err != nil {
		session.Send(channelID, "エラーが発生しました")
		return err
	}

	return err
}
