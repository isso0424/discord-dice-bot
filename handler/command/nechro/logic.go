package nechro

import (
	"fmt"
	"isso0424/dise/dice"
)

const template = "修正値: %d\nダイス: %d\n判定値: %d\n結果: %s"

func calcJudgeResult(amendmentPoint int) int {
	d := dice.New()

	return d.RollOne(10) + amendmentPoint
}

func judge(amendmentPoint int) (string, int, bool) {
	result := calcJudgeResult(amendmentPoint)
	r := getResultComment(result >= 6)

	return fmt.Sprintf(template, amendmentPoint, result-amendmentPoint, result, r), result, result >= 6
}

func getResultComment(result bool) string {
	if result {
		return "成功"
	}

	return "失敗"
}
