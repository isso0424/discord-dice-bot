package parser_test

import (
	"isso0424/dice/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		t.Fatalf("should occur error on here")

		return
	}

	assert.Equal(t, "command not provides", err.Error())
}
