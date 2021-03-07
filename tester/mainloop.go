package tester

import (
	"fmt"
	"isso0424/dise/handler/console"
)

type mainloop struct {
	channelID string
}

func New(channelID string) mainloop {
	return mainloop{ channelID }
}

func(loop mainloop) Start() {
	executer := console.New()
	for {
		fmt.Print("input text\n>>>")
		text := receive()

		executer.Send(loop.channelID, text)
	}
}
