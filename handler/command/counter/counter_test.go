package counter_test

import (
	"isso0424/dise/handler/command/counter"
	"isso0424/dise/handler/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldErrorOccurInHere = "should error occur with "

var (
	c = counter.Counter{}
	s = console.New()
)

func TestCounterExecSuccess(t *testing.T) {
	err := c.Exec("hoge", []string{"60", "30"}, s)
	if err != nil {
		t.Fatal(err)

		return
	}
}

func TestCounterExecFail(t *testing.T) {
	err := c.Exec("", []string{"60", "30"}, s)
	if err == nil {
		t.Fatal(shouldErrorOccurInHere, "invalid channel id")

		return
	}

	err = c.Exec("hoge", []string{}, s)
	if err == nil {
		t.Fatal(shouldErrorOccurInHere, "empty args")

		return
	}

	err = c.Exec("hoge", []string{"invalid", "args"}, s)
	if err == nil {
		t.Fatal(shouldErrorOccurInHere, "invalid args")

		return
	}
}

func TestCounterGetHelp(t *testing.T) {
	assert.Equal(t, "対抗ロールを行います", c.GetHelp())
}

func TestCounterGetPrefix(t *testing.T) {
	assert.Equal(t, "!counter", c.GetPrefix())
}

func TestCounterValidateArgs(t *testing.T) {
	assert.Equal(t, true, c.ValidateArgs([]string{"60", "30"}))
	assert.Equal(t, false, c.ValidateArgs([]string{"60", "invalid"}))
	assert.Equal(t, false, c.ValidateArgs([]string{"invalid", "60"}))
	assert.Equal(t, false, c.ValidateArgs([]string{}))
}
