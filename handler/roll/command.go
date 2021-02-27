package roll

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Roll(channelID string, args []string, session *discordgo.Session) {
	results, err := allRoll(args)
	if err != nil {
		_, err = session.ChannelMessageSend(channelID, err.Error())
		if err != nil {
			fmt.Println(err)

			return
		}

		return
	}

	message := joinResults(results)

	_, err = session.ChannelMessageSend(channelID, message)
	if err != nil {
		fmt.Println(err)

		return
	}
	fmt.Println(message)
}
