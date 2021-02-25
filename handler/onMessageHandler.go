package handler

import (
	"isso0424/dice/dice"
	"isso0424/dice/handler/judge"
	"isso0424/dice/handler/roll"
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
		roll.Roll(event.ChannelID, args, d, session)
	}

	if command == "!judge" && len(args) != 0 {
		judge.Judge(event.ChannelID, args, d, session)
	}
}
