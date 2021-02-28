package dice_test

import (
	"isso0424/dise/dice"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiceRoll(t *testing.T) {
	d := dice.New()

	result := d.Roll(6, 100)
	for _, r := range result {
		if r > 6 || r < 1 {
			t.Fatal("Broken dice")

			return
		}
	}

	assert.Equal(t, 100, len(result))
}

func TestDiceRollOne(t *testing.T) {
	d := dice.New()

	for i := 0; i < 100; i++ {
		result := d.RollOne(3)
		if result > 3 || result < 1 {
			t.Fatal("Broken dice")

			return
		}
	}
}
