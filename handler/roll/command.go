package roll

import (
	"fmt"
	"isso0424/dise/types"
)

func Roll(channelID string, args []string, session types.Session) error {
	results, err := allRoll(args)
	if err != nil {
		return err
	}

	message := joinResults(results)

	err = session.Send(channelID, message)
	if err != nil {
		return err
	}
	fmt.Println(message)

	return nil
}
