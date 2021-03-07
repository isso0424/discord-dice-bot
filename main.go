package main

import (
	"bufio"
	"flag"
	"fmt"
	"isso0424/dise/handler"
	"isso0424/dise/tester"
	"isso0424/dise/tool"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type cmdArgs struct {
	IsTestMode bool
	IsReadmeUpdateFlag bool
}

func parseArgs() cmdArgs {
	isTestFlag := flag.Bool("test", false, "command line test mode")
	isReadmeUpdateFlag := flag.Bool("update-readme", false, "update readme with command help")
	flag.Parse()

	return cmdArgs{ IsTestMode: *isTestFlag, IsReadmeUpdateFlag: *isReadmeUpdateFlag }
}

func main() {
	args := parseArgs()

	if args.IsReadmeUpdateFlag {
		tool.UpdateReadme()
		log.Println("done")
		return
	}

	if args.IsTestMode {
		fmt.Print("Please input channel id\n>>>")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()

		tester.New(stdin.Text()).Start()
		return
	}
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

	log.Println("bot booted!!!")

	defer discord.Close()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-sc
}
