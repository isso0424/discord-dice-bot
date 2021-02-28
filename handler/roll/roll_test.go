package roll_test

import (
	"fmt"
	"isso0424/dise/handler/console"
	"isso0424/dise/handler/roll"
	"testing"
)

const shouldErrorOccur = "should error occur with %s"

func TestRollSuccess(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	dices := []string{"3D6", "1D10", "2D2"}

	err := roll.Roll(channelID, dices, session)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRollFail(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	dices := []string{"3D6", "1D10", "2D2"}

	invalidChannelID := ""
	err := roll.Roll(invalidChannelID, dices, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "empty channelID"))
	}

	invalidArgs := []string{"invalid"}
	err = roll.Roll(channelID, invalidArgs, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "invalid args"))
	}
}
