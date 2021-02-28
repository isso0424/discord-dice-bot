package main

import (
	"fmt"
	"isso0424/dise/handler"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New()
	discord.Token = fmt.Sprintf("Bot %s", os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	discord.AddHandler(handler.OnMessageHandler)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	defer discord.Close()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}
