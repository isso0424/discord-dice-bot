package tool

import (
	"fmt"
	"isso0424/dise/handler/command"
	"log"
	"os"
)

const (
	readmeTemplate = `# dise
[![Build](https://github.com/isso0424/dise/actions/workflows/build.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/build.yml)
[![Lint](https://github.com/isso0424/dise/actions/workflows/lint.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/lint.yml)
[![Test](https://github.com/isso0424/dise/actions/workflows/test.yml/badge.svg)](https://github.com/isso0424/dise/actions/workflows/test.yml)
[![Issue](https://img.shields.io/github/issues/isso0424/dise)](https://github.com/isso0424/dise/issues)
![Language](https://img.shields.io/github/languages/top/isso0424/dise)
![License](https://img.shields.io/github/license/isso0424/dise)
## About
Dice bot for discord.

## Usage
`
	usageTemplate = `### %s
%s

`
)

func UpdateReadme() {
	usageText := ""
	for _, cmd := range command.Commands {
		usageText += fmt.Sprintf(usageTemplate, cmd.GetPrefix(), cmd.GetHelp())
	}

	f, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.Write([]byte(readmeTemplate + usageText))
	if err != nil {
		log.Fatal(err)
	}
}
