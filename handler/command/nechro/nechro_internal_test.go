package nechro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcJudgeResult(t *testing.T) {
	result := calcJudgeResult(10)
	assert.Equal(t, true, result > 10)
}

func TestJudge(t *testing.T) {
	_, result, r := judge(5)
	assert.Equal(t, true, result >= 6)
	assert.Equal(t, true, r)

	_, result, r = judge(-5)
	assert.Equal(t, false, result >= 6)
	assert.Equal(t, false, r)
}
