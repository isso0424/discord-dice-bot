package handler

import (
	"isso0424/dise/handler/command"
	"isso0424/dise/handler/discord"
	"isso0424/dise/handler/judge"
	"isso0424/dise/handler/roll"
	"isso0424/dise/parser"
	"log"

	"github.com/bwmarrin/discordgo"
)

const errorMessage = "エラーが発生しました"

var (
	judgeCommand = judge.Judge{}
	rollCommand = roll.Roll{}
)

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	s := discord.CreateSenderFromSession(session)
	channelID := event.ChannelID

	cmd, args, err := parser.ParseCommand(event.Content)
	if err != nil {
		_, err = session.ChannelMessageSend(event.ChannelID, "command parsing error")
		if err != nil {
			log.Println(err)
		}

		return
	}

	switch cmd {
	case judgeCommand.GetPrefix():
		err := command.Exec(judgeCommand, channelID, args, s)
		log.Println(err)
	case rollCommand.GetPrefix():
		err := command.Exec(rollCommand, channelID, args, s)
		log.Println(err)
	}
}
