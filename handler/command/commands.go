package command

import (
	"isso0424/dise/handler/command/counter"
	"isso0424/dise/handler/command/judge"
	"isso0424/dise/handler/command/roll"
)

var commands = []Command{
	counter.Counter{},
	judge.Judge{},
	roll.Roll{},
	help{},
}
