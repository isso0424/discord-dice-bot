package roll

import (
	"isso0424/dise/types"
	"log"
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
	log.Println(message)

	return nil
}
