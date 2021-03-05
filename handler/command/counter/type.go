package counter

type diceResult int

const (
	success = iota
	fail
	autoSuccess
	autoFailed
)

func (r diceResult) String() string {
	switch r {
	case success:
		return "成功"
	case fail:
		return "失敗"
	case autoSuccess:
		return "自動成功"
	case autoFailed:
		return "自動失敗"
	default:
		return "不正な結果です"
	}
}
