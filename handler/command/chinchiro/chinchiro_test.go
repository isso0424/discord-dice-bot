package chinchiro_test

import (
	"isso0424/dise/handler/command/chinchiro"
	"isso0424/dise/handler/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldErrorOccurInHere = "should error occur with "

var (
	cmd     = chinchiro.Chinchiro{}
	session = console.New()
)

func TestChinchiroExecSuccess(t *testing.T) {
	err := cmd.Exec("hoge", []string{}, session)
	if err != nil {
		t.Fatal(err)
	}

	err = cmd.Exec("hoge", []string{"args"}, session)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChinchiroExecFail(t *testing.T) {
	err := cmd.Exec("", []string{}, session)
	if err == nil {
		t.Fatal(shouldErrorOccurInHere, "invalid channel")
	}
}

func TestGetPrefix(t *testing.T) {
	assert.Equal(t, "!chinchiro", cmd.GetPrefix())
}

func TestGetHelp(t *testing.T) {
	assert.Equal(t, "チンチロをします。", cmd.GetHelp())
}

func TestValidateArgs(t *testing.T) {
	assert.Equal(t, true, cmd.ValidateArgs([]string{}))
}
