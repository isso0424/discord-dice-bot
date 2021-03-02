package command_test

import (
	"github.com/stretchr/testify/assert"
	"isso0424/dise/handler/command"
	"isso0424/dise/handler/console"
	"testing"
)

var session = console.New()

func TestExecuteCommandSuccess(t *testing.T) {
	err := command.ExecuteCommand("!judge 80", "hoge", session)
	assert.Equal(t, nil, err)

	err = command.ExecuteCommand("!help", "fuga", session)
	assert.Equal(t, nil, err)

	err = command.ExecuteCommand("!roll 1D3", "foo", session)
	assert.Equal(t, nil, err)

	err = command.ExecuteCommand("!invalid", "bar", session)
	assert.Equal(t, nil, err)
}

func TestExecuteCommandFail(t *testing.T) {
	err := command.ExecuteCommand("", "hoge", session)
	if err == nil {
		t.Fatal("Should error occur with empty message")
	}

	err = command.ExecuteCommand("!judge hoge", "hoge", session)
	if err == nil {
		t.Fatal("Should error occur with command failed")
	}

	err = command.ExecuteCommand("!judge 80", "", session)
	if err == nil {
		t.Fatal("Should error occur with empty channel ID")
	}

	err = command.ExecuteCommand("", "", session)
	if err == nil {
		t.Fatal("Should error occur with empty channel ID")
	}
}
