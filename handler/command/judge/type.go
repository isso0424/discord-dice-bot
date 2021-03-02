package judge

type DiceResult int

const (
	Success = iota
	Fail
	Critical
	Fumble
)

func (r DiceResult) String() string {
	switch r {
	case Success:
		return "成功"
	case Fail:
		return "失敗"
	case Critical:
		return "クリティカル"
	case Fumble:
		return "ファンブル"
	default:
		return "異常な値です"
	}
}
