package roll

import (
	"fmt"
	"isso0424/dice/dice"
	"isso0424/dice/parser"

	"github.com/bwmarrin/discordgo"
)

func Roll(channelID string, args []string, d dice.Dice, session *discordgo.Session) {
	resultSum := 0
	resultString := "["

	for _, command := range args {
		count, max, err := parser.ParseDice(command)
		if err != nil {
			session.ChannelMessageSend(channelID, err.Error())

			return
		}

		if (max < 1 || count < 1) {
			session.ChannelMessageSend(channelID, "値は1以上でなければいけません")

			return
		}

		result := d.Roll(max, count)
		for _, value := range result {
			resultSum += value
			resultString += fmt.Sprintf("%d,", value)
		}
	}

	resultString += "]"

	_, err := session.ChannelMessageSend(channelID, fmt.Sprintf("合計: %d %s", resultSum, resultString))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultString)
}
