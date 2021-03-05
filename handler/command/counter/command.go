package counter

import (
	"fmt"
	"isso0424/dise/types"
	"strconv"
)

type Counter struct {
}

func (cmd Counter) Exec(channelID string, args []string, session types.Session) error {
	action, reaction, err := parseArgs(args)
	if err != nil {
		return err
	}

	result, dice := judge(action, reaction)
	return session.Send(channelID, fmt.Sprintf("結果: %s\n出目: %d", result.String(), dice))
}

func (cmd Counter) ValidateArgs(args []string) bool {
	if len(args) < 2 {
		return false
	}

	if n, err := strconv.Atoi(args[0]); err != nil || n < 0 {
		return false
	}

	if n, err := strconv.Atoi(args[1]); err != nil || n < 0 {
		return false
	}

	return true
}

func (cmd Counter) GetPrefix() string {
	return "!counter"
}

func (cmd Counter) GetHelp() string {
	return "対抗ロールを行います。自動成功/失敗は出目が100で固定です。"
}
