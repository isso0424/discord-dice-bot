package discord

import "github.com/bwmarrin/discordgo"

type session struct {
	session *discordgo.Session
}

func (s session) Send(channelID string, message string) error {
	_, err := s.session.ChannelMessageSend(channelID, message)

	return err
}

func CreateSenderFromSession(s *discordgo.Session) (session) {
	return session{ session: s }
}
