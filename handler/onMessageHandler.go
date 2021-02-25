package handler

import (
	"errors"
	"fmt"
	"isso0424/dice/dice"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var d = dice.Dice{}

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(event.Content, "!roll") && len(strings.Split(event.Content, " ")) > 1 {
		commands := strings.Split(event.Content, " ")[1:]

		resultSum := 0
		resultString := "["

		for _, command := range commands {
			max, count, err := diceRollParsing(command)
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

func diceRollParsing(text string) (int, int, error) {
	command := strings.Split(text, "D")
	if len(command) < 2 {
		return 0, 0, errors.New("invalid command")
	}

	count, err := strconv.Atoi(command[0])
	if err != nil {
		return 0, 0, errors.New("parsing error")
	}

	max, err := strconv.Atoi(command[1])
	if err != nil {
		return 0, 0, errors.New("parsing error")
	}

	return max, count, nil
}
