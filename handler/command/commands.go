package command

import (
	"isso0424/dise/handler/command/judge"
	"isso0424/dise/handler/command/roll"
)

var (
	judgeCommand = judge.Judge{}
	rollCommand  = roll.Roll{}
	commands     = map[string]Command{
		judgeCommand.GetPrefix(): judgeCommand,
		rollCommand.GetPrefix():  rollCommand,
	}
)
