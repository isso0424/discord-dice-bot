package command

import "isso0424/dise/types"

type Command interface {
	GetHelp() string
	GetPrefix() string
	ValidateArgs(args []string) bool
	Exec(channelID string, args []string, session types.Session) error
}
