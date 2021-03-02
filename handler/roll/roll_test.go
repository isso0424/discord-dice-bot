package roll_test

import (
	"fmt"
	"isso0424/dise/handler/console"
	"isso0424/dise/handler/roll"
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldErrorOccur = "should error occur with %s"

var rollCommand = roll.Roll{}

func TestRollSuccess(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	dices := []string{"3D6", "1D10", "2D2"}

	err := rollCommand.Exec(channelID, dices, session)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRollFail(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	dices := []string{"3D6", "1D10", "2D2"}

	invalidChannelID := ""
	err := rollCommand.Exec(invalidChannelID, dices, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "empty channelID"))
	}

	invalidArgs := []string{"invalid"}
	err = rollCommand.Exec(channelID, invalidArgs, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "invalid args"))
	}
}

func TestValidateArgsSuccess(t *testing.T) {
	args := []string{ "1D2", "3D6" }
	assert.Equal(t, true, rollCommand.ValidateArgs(args))
}

func TestValidateArgsFail(t *testing.T) {
	args := []string{"invalid"}
	assert.Equal(t, false, rollCommand.ValidateArgs(args))

	args = []string{}
	assert.Equal(t, false, rollCommand.ValidateArgs(args))
}

func TestGetPrefix(t *testing.T) {
	assert.Equal(t, "!roll", rollCommand.GetPrefix())
}
