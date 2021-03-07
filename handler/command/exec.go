package command

import (
	"fmt"
	"isso0424/dise/parser"
	"isso0424/dise/types"
	"log"
)

func ExecuteCommand(text string, channelID string, session types.Session) error {
	cmd, args, err := parser.ParseCommand(text)
	if err != nil {
		e := session.Send(channelID, "command parsing error")
		if e != nil {
			log.Println(e)
			return err
		}

		return err
	}

	for _, command := range Commands {
		if command.GetPrefix() == cmd {
			err := exec(command, channelID, args, session)

			return err
		}
	}

	return nil
}

func exec(cmd Command, channelID string, args []string, session types.Session) error {
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
