package parser_test

import (
	"isso0424/dise/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldOccurErrorMessage = "should occur with %s"

func TestParseCommandOnly(t *testing.T) {
	command := "!hoge"
	result, args, err := parser.ParseCommand(command)
	if err != nil {
		t.Fatalf(err.Error())

		return
	}

	if len(args) != 0 {
		t.Fatalf("Invalid args")

		return
	}

	assert.Equal(t, "!hoge", result)
}

func TestParseCommandWithArgs(t *testing.T) {
	command := "!hoge fuga foo bar"
	result, args, err := parser.ParseCommand(command)
	if err != nil {
		t.Fatalf(err.Error())

		return
	}

	if len(args) != 3 {
		t.Fatalf("Invalid args")

		return
	}

	assert.Equal(t, "!hoge", result)
	assert.Equal(t, "fuga", args[0])
	assert.Equal(t, "foo", args[1])
	assert.Equal(t, "bar", args[2])
}

func TestParseEmptyCommand(t *testing.T) {
	command := ""
	_, _, err := parser.ParseCommand(command)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "empty command")

		return
	}

	assert.Equal(t, "command not provides", err.Error())
}

func TestDiceParseSuccess(t *testing.T) {
	dice := "3D6"
	count, max, err := parser.ParseDice(dice)
	if err != nil {
		t.Fatalf(err.Error())

		return
	}

	assert.Equal(t, 3, count)
	assert.Equal(t, 6, max)
}

func TestDiceParseFail(t *testing.T) {
	notDice := "123"
	_, _, err := parser.ParseDice(notDice)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "dice: 123")

		return
	}
	assert.Equal(t, "invalid dice", err.Error())

	abnormalDice := "1D"
	_, _, err = parser.ParseDice(abnormalDice)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "dice: 1D")

		return
	}
	assert.Equal(t, "parsing error", err.Error())

	niceCountDice := "invalidD6"
	_, _, err = parser.ParseDice(niceCountDice)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "dice: invalidD6")

		return
	}
	assert.Equal(t, "parsing error", err.Error())

	niceMaxDice := "1Dinvalid"
	_, _, err = parser.ParseDice(niceMaxDice)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "dice: 1Dinvalid")

		return
	}
	assert.Equal(t, "parsing error", err.Error())

	chainingDice := "1D6D6D3"
	_, _, err = parser.ParseDice(chainingDice)
	if err == nil {
		t.Fatalf(shouldOccurErrorMessage, "dice: 1Dinvalid")

		return
	}
	assert.Equal(t, "invalid dice", err.Error())
}
