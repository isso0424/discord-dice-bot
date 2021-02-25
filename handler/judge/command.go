package judge

import (
	"fmt"
	"isso0424/dice/dice"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Judge(channelID string, args []string, d dice.Dice, session *discordgo.Session) {
	if len(args) == 0 {
		session.ChannelMessageSend(channelID, "too few arguments")

		return
	}

	target, err := strconv.Atoi(args[0])
	if err != nil {
		session.ChannelMessageSend(channelID, "invalid parsing")

		return
	}

	result := d.Roll(100, 1)[0]

	var status string
	if result <= target {
		status = "成功!"
	} else {
		status = "失敗"
	}

	session.ChannelMessageSend(channelID, fmt.Sprintf("目標値: %d\nダイス: %d\n結果: %s", target, result, status))
}
