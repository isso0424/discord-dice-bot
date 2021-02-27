package judge

import (
	"fmt"
	"isso0424/dice/dice"

	"github.com/bwmarrin/discordgo"
)

func Judge(channelID string, args []string, d dice.Dice, session *discordgo.Session) {
	target, err := validateArgs(args)
	if err != nil {
		session.ChannelMessageSend(channelID, err.Error())

		return
	}

	diceResult := d.Roll(100, 1)[0]
	result := compareResult(target, diceResult)

	session.ChannelMessageSend(channelID, fmt.Sprintf("目標値: %d\nダイス: %d\n結果: %s", target, diceResult, result.String()))
}
