package roll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldOccurErrorMessage = "should occur error with %s"

func TestAllRollSuccess(t *testing.T) {
	dices := []string{"3D6", "2D3"}
	results, err := allRoll(dices)
	if err != nil {
		t.Fatal(err.Error())

		return
	}

	assert.Equal(t, 5, len(results))
}

func TestAllRollFail(t *testing.T) {
	dices := []string{"3D"}
	_, err := allRoll(dices)
	if err == nil {
		t.Fatal(err.Error())

		return
	}
}

func TestJoinResults(t *testing.T) {
	args := []int{0, 1, 2, 3}
	result := joinResults(args)

	assert.Equal(t, "合計: 6 [0,1,2,3]", result)
}
