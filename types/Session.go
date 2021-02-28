package types

type Session interface {
	ChannelMessageSend(string, string) error
}
