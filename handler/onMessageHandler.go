package handler

import (
	"fmt"
	"isso0424/dise/handler/judge"
	"isso0424/dise/handler/roll"
	"isso0424/dise/parser"

	"github.com/bwmarrin/discordgo"
)

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	command, args, err := parser.ParseCommand(event.Content)
	if err != nil {
		_, err = session.ChannelMessageSend(event.ChannelID, "command parsing error")
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	if command == "!roll" && len(args) != 0 {
		roll.Roll(event.ChannelID, args, session)
	}

	if command == "!judge" && len(args) != 0 {
		judge.Judge(event.ChannelID, args, session)
	}
}
