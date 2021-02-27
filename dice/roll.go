package dice

import "math/rand"

func (dice Dice) Roll(max int, count int) (result []int) {
	for i := 0; i < count; i++ {
		result = append(result, rand.Intn(max)+1)
	}

	return result
}

func (dice Dice) RollOne(max int) (result int) {
	return rand.Intn(max) + 1
}
