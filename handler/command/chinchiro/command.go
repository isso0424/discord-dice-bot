package chinchiro

import "isso0424/dise/types"

type Chinchiro struct {
}

func (cmd Chinchiro) Exec(channelID string, args []string, session types.Session) error {
	return session.Send(channelID, roll())
}

func (cmd Chinchiro) GetPrefix() string {
	return "!chinchiro"
}

func (cmd Chinchiro) GetHelp() string {
	return "チンチロをします。"
}

func (cmd Chinchiro) ValidateArgs(args []string) bool {
	return true
}
