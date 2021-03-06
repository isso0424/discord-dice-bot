package chinchiro

type chinchiroResult int

const (
	pinZoro = iota
	niZoro
	sanZoro
	yonZoro
	goZoro
	rokuZoro
	hihumi
	shigoro
	demeIchi
	demeNi
	demeSan
	demeYon
	demeGo
	demeRoku
	yakuNashi
)

func(result chinchiroResult) String() string {
	switch result {
	case pinZoro:
		return "ピンゾロ"
	case niZoro:
		return "ニゾロ"
	case sanZoro:
		return "サンゾロ"
	case yonZoro:
		return "ヨンゾロ"
	case goZoro:
		return "ゴゾロ"
	case rokuZoro:
		return "ロクゾロ"
	case hihumi:
		return "ヒフミ"
	case shigoro:
		return "シゴロ"
	case demeIchi:
		return "出目イチ"
	case demeNi:
		return "出目ニ"
	case demeSan:
		return "出目サン"
	case demeYon:
		return "出目ヨン"
	case demeGo:
		return "出目ゴ"
	case demeRoku:
		return "出目ロク"
	case yakuNashi:
		return "役無し"
	default:
		return "エラー"
	}
}
