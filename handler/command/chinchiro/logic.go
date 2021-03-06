package chinchiro

import (
	"fmt"
	"isso0424/dise/dice"
	"sort"
)

func roll() string {
	d := dice.Dice{}

	result := d.Roll(6, 3)

	return fmt.Sprintf("役: %s\n出目: %v", judge(result).String(), result)
}

func judge(result []int) chinchiroResult {
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	if result[0] == result[1] && result[1] == result[2] {
		return zoro(result[0])
	}

	if result[0] == result[1] || result[1] == result[2] || result[2] == result[0] {
		return deme(result)
	}

	if result[0] == 1 && result[1] == 2 && result[2] == 3 {
		return hihumi
	}

	if result[0] == 4 && result[1] == 5 && result[2] == 6 {
		return shigoro
	}

	return yakuNashi
}

func zoro(diceResult int) chinchiroResult {
	switch diceResult {
	case 1:
		return pinZoro
	case 2:
		return niZoro
	case 3:
		return sanZoro
	case 4:
		return yonZoro
	case 5:
		return goZoro
	case 6:
		return rokuZoro
	default:
		return yakuNashi
	}
}

func deme(diceResults []int) chinchiroResult {
	var resultPosition int
	if diceResults[0] == diceResults[1] {
		resultPosition = 2
	}
	if diceResults[1] == diceResults[2] {
		resultPosition = 0
	}
	if diceResults[2] == diceResults[0] {
		resultPosition = 1
	}

	switch diceResults[resultPosition] {
	case 1:
		return demeIchi
	case 2:
		return demeNi
	case 3:
		return demeSan
	case 4:
		return demeYon
	case 5:
		return demeGo
	case 6:
		return demeRoku
	default:
		return yakuNashi
	}
}
