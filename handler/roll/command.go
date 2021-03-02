package roll

import (
	"isso0424/dise/parser"
	"isso0424/dise/types"
	"log"
)

type Roll struct {}

func(cmd Roll) Exec(channelID string, args []string, session types.Session) error {
	results, err := allRoll(args)
	if err != nil {
		return err
	}

	message := joinResults(results)

	err = session.Send(channelID, message)
	if err != nil {
		return err
	}
	log.Println(message)

	return nil
}

func(cmd Roll) ValidateArgs(args []string) bool {
	if len(args) == 0 {
		return false
	}

	for _, arg := range args {
		_, _, err := parser.ParseDice(arg)
		if err != nil {
			return false
		}
	}

	return true
}

func(cmd Roll) GetPrefix() string {
	return "!roll"
}

func(cmd Roll) GetHelp() string {
	return "指定された方法でダイスを振り、合計を出力します。"
}
