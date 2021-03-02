package command

import (
	"isso0424/dise/handler/judge"
	"isso0424/dise/handler/roll"
)

var (
	judgeCommand = judge.Judge{}
	rollCommand = roll.Roll{}
  commands = map[string]Command{
		judgeCommand.GetPrefix(): judgeCommand,
		rollCommand.GetPrefix(): rollCommand,
	}
)
