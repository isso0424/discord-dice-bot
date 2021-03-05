package counter

import (
	"isso0424/dise/dice"
	"strconv"
)

func parseArgs(args []string) (active int, reaction int, err error) {
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

func judge(action, reaction int) diceResult {
	d := dice.Dice{}

	target := action - reaction + 50
	if target > 100 {
		return autoSuccess
	}
	if target < 0 {
		return autoFailed
	}

	result := d.RollOne(100)
	if target >= result {
		return fail
	}

	return success
}
