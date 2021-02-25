package handler

import (
	"fmt"
	"isso0424/dice/dice"
	"isso0424/dice/parser"

	"github.com/bwmarrin/discordgo"
)

var d = dice.Dice{}

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	command, args, err := parser.ParseCommand(event.Content)
	if err != nil {
		session.ChannelMessageSend(event.ChannelID, "command parsing error")

		return
	}

	if command == "!roll" && len(args) != 0 {

		resultSum := 0
		resultString := "["

		for _, command := range args {
			count, max, err := parser.ParseDice(command)
			if err != nil {
				session.ChannelMessageSend(event.ChannelID, err.Error())

				return
			}

			if (max < 1 || count < 1) {
				session.ChannelMessageSend(event.ChannelID, "値は1以上でなければいけません")

				return
			}

			result := d.Roll(max, count)
			for _, value := range result {
				resultSum += value
				resultString += fmt.Sprintf("%d,", value)
			}
		}

		resultString += "]"

		_, err := session.ChannelMessageSend(event.ChannelID, fmt.Sprintf("合計: %d %s", resultSum, resultString))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resultString)
	}
}
