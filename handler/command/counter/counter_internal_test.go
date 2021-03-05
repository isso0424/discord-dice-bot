package counter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudge(t *testing.T) {
	result, _ := judge(100, 0)
	assert.Equal(t, diceResult(autoSuccess), result)

	result, _ = judge(0, 100)
	assert.Equal(t, diceResult(autoFailed), result)

	result, _ = judge(50, 0)
	assert.Equal(t, diceResult(success), result)

	result, _ = judge(0, 50)
	assert.Equal(t, diceResult(fail), result)

	result, _ = judge(50, 50)
	assert.Equal(t, true, result == success || result == fail)
}

func TestParseArgsSuccess(t *testing.T) {
	correctArgs := []string{"60", "30"}

	active, reaction, err := parseArgs(correctArgs)
	assert.Equal(t, 60, active)
	assert.Equal(t, 30, reaction)
	assert.Equal(t, nil, err)
}

func TestParseArgsFail(t *testing.T) {
	invalidArgs := []string{}

	_, _, err := parseArgs(invalidArgs)
	if err == nil {
		t.Fatal(err)

		return
	}

	_, _, err = parseArgs([]string{"invalid", "60"})
	if err == nil {
		t.Fatal(err)

		return
	}

	_, _, err = parseArgs([]string{"60", "invalid"})
	if err == nil {
		t.Fatal(err)

		return
	}
}

func TestDiceResult(t *testing.T) {
	assert.Equal(t, "成功", diceResult(success).String())
	assert.Equal(t, "失敗", diceResult(fail).String())
	assert.Equal(t, "自動成功", diceResult(autoSuccess).String())
	assert.Equal(t, "自動失敗", diceResult(autoFailed).String())
	assert.Equal(t, "不正な結果です", diceResult(5).String())
}
