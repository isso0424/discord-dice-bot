package roll

import (
	"fmt"
	"strconv"

	"isso0424/dice/dice"
	"isso0424/dice/parser"
)

const resultFormat = "合計: %d [%s]"

func allRoll(args []string) (results []int, err error) {
	d := dice.New()
	for index, arg := range args {
		count, max, err := parser.ParseDice(arg)
		if err != nil {
			return nil, fmt.Errorf("%s: on %d dice", err, index+1)
		}
		result := d.Roll(max, count)
		results = append(results, result...)
	}

	return
}

func joinResults(results []int) string {
	resultSum := 0
	resultString := ""

	for index, result := range results {
		resultSum += result
		resultString += strconv.Itoa(result)
		if index+1 != len(results) {
			resultString += ","
		}
	}

	return fmt.Sprintf(resultFormat, resultSum, resultString)
}
