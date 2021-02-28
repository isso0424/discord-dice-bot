package roll

import (
	"fmt"
	"isso0424/dise/types"
)

func Roll(channelID string, args []string, session types.Session) {
	results, err := allRoll(args)
	if err != nil {
		err = session.Send(channelID, err.Error())
		if err != nil {
			fmt.Println(err)

			return
		}

		return
	}

	message := joinResults(results)

	err = session.Send(channelID, message)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(message)
}
