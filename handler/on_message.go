package handler

import (
	"isso0424/dise/handler/command"
	"isso0424/dise/handler/discord"

	"github.com/bwmarrin/discordgo"
)

func OnMessageHandler(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return
	}

	s := discord.CreateSenderFromSession(session)

	command.ExecuteCommand(event.Content, event.ChannelID, s)
}
