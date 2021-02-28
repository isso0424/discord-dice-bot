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
