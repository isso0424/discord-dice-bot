package judge

import (
	"errors"
	"strconv"
)

func compareResult(target int, dice int) (result DiceResult) {
	switch {
	case dice <= target && dice <= 5:
		return Critical
	case dice <= target:
		return Success
	case dice >= 96:
		return Fumble
	default:
		return Fail
	}
}

func validateArgs(args []string) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("too few arguments")
	}

	target, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.New("failed parsing")
	}

	return target, nil
}
