package parser

import (
	"errors"
	"strconv"
	"strings"
)

func ParseDice(text string) (count int, max int, err error) {
	if !strings.Contains(text, "D") {
		err = errors.New("invalid dice")

		return
	}

	texts := strings.Split(text, "D")
	if len(texts) < 2 {
		err = errors.New("invalid dice")

		return
	}

	count, err = strconv.Atoi(texts[0])
	if err != nil {
		err = errors.New("parsing error")

		return
	}

	max, err = strconv.Atoi(texts[1])
	if err != nil {
		err = errors.New("parsing error")

		return
	}

	return
}
