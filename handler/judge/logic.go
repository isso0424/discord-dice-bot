package judge

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
