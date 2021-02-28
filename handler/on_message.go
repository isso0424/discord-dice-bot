package handler

import (
	"isso0424/dise/handler/discord"
	"isso0424/dise/handler/judge"
	"isso0424/dise/handler/roll"
	"isso0424/dise/parser"
	"log"

	"github.com/bwmarrin/discordgo"
)

const errorMessage = "エラーが発生しました"

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	s := discord.CreateSenderFromSession(session)

	command, args, err := parser.ParseCommand(event.Content)
	if err != nil {
		_, err = session.ChannelMessageSend(event.ChannelID, "command parsing error")
		if err != nil {
			log.Println(err)
		}

		return
	}

	if command == "!roll" && len(args) != 0 {
		err = roll.Roll(event.ChannelID, args, s)
		if err != nil {
			s.Send(event.ChannelID, errorMessage)
			log.Println(err)
		}
	}

	if command == "!judge" && len(args) != 0 {
		err = judge.Judge(event.ChannelID, args, s)
		if err != nil {
			s.Send(event.ChannelID, errorMessage)
			log.Println(err)
		}
	}
}
