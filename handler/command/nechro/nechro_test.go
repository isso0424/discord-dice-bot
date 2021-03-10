package nechro_test

import (
	"isso0424/dise/handler/command/nechro"
	"isso0424/dise/handler/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	command = nechro.Nechro{}
	session = console.New()
)

func TestExecCommandSuccess(t *testing.T) {
	err := command.Exec("hoge", []string{"-1"}, session)
	assert.Equal(t, nil, err)

	err = command.Exec("hoge", []string{"1"}, session)
	assert.Equal(t, nil, err)
}

func TestExecCommandFail(t *testing.T) {
	err := command.Exec("", []string{"1"}, session)
	if err == nil {
		shouldErrorOccur(t, "empty channel id")
	}

	err = command.Exec("hoge", []string{"a"}, session)
	if err == nil {
		shouldErrorOccur(t, "string args")
	}

	err = command.Exec("hoge", []string{}, session)
	if err == nil {
		shouldErrorOccur(t, "empty args")
	}
}

func TestGetPrefix(t *testing.T) {
	assert.Equal(t, "!nechro", command.GetPrefix())
}

func TestGetHelp(t *testing.T) {
	assert.Equal(t, "修正値を基にネクロニカの行動判定をします。", command.GetHelp())
}

func TestValidateArgs(t *testing.T) {
	result := command.ValidateArgs([]string{"1"})
	assert.Equal(t, true, result)

	result = command.ValidateArgs([]string{"invalid"})
	assert.Equal(t, false, result)

	result = command.ValidateArgs([]string{})
	assert.Equal(t, false, result)
}

func shouldErrorOccur(t *testing.T, message string) {
	t.Fatal("should error occur with ", message)
}
