package judge

import (
	"fmt"
	"isso0424/dise/dice"
	"isso0424/dise/types"
)

func Judge(channelID string, args []string, session types.Session) error {
	d := dice.New()

	target, err := validateArgs(args)
	if err != nil {
		return err
	}

	diceResult := d.RollOne(100)
	result := compareResult(target, diceResult)

	err = session.Send(channelID, fmt.Sprintf("目標値: %d\nダイス: %d\n結果: %s", target, diceResult, result.String()))

	return err
}
