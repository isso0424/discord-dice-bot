package judge

import (
	"fmt"
	"isso0424/dise/dice"
	"isso0424/dise/types"
	"strconv"
)

type Judge struct {
}

func(c Judge) Exec(channelID string, args []string, session types.Session) error {
	d := dice.New()

	target, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	diceResult := d.RollOne(100)
	result := compareResult(target, diceResult)

	err = session.Send(channelID, fmt.Sprintf("目標値: %d\nダイス: %d\n結果: %s", target, diceResult, result.String()))

	return err
}

func(c Judge) ValidateArgs(args []string) bool {
	if len(args) == 0 {
		return false
	}

	_, err := strconv.Atoi(args[0])
	if err != nil {
		return false
	}

	return true
}

func(c Judge) GetPrefix() string {
	return "!judge"
}
