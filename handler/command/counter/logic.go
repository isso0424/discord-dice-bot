package counter

import (
	"errors"
	"isso0424/dise/dice"
	"strconv"
)

func parseArgs(args []string) (active int, reaction int, err error) {
	if len(args) < 2 {
		err = errors.New("Invalid length")
		return
	}

	active, err = strconv.Atoi(args[0])
	if err != nil {
		return
	}

	reaction, err = strconv.Atoi(args[1])
	if err != nil {
		return
	}

	return
}

func judge(action, reaction int) (diceResult, int) {
	d := dice.Dice{}

	target := action - reaction + 50
	if target > 100 {
		return autoSuccess, 100
	}
	if target < 0 {
		return autoFailed, 100
	}

	result := d.RollOne(100)
	if result >= target {
		return fail, result
	}

	return success, result
}
