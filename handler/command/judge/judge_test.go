package judge_test

import (
	"fmt"
	"isso0424/dise/handler/command/judge"
	"isso0424/dise/handler/console"
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldErrorOccur = "should error occur with %s"

var judgeCommand = judge.Judge{}

func TestJudgeSuccess(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	args := []string{"100"}

	err := judgeCommand.Exec(channelID, args, session)
	if err != nil {
		t.Fatal(err)
	}
}

func TestJudgeFail(t *testing.T) {
	session := console.New()
	channelID := "hoge"
	args := []string{"100"}

	invalidChannelID := ""
	err := judgeCommand.Exec(invalidChannelID, args, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "invalid channel ID"))
	}

	invalidArgs := []string{"invalid"}
	err = judgeCommand.Exec(channelID, invalidArgs, session)
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "invalid args"))
	}
}

func TestDiceResult(t *testing.T) {
	success := judge.DiceResult(0)
	fail := judge.DiceResult(1)
	critical := judge.DiceResult(2)
	fumble := judge.DiceResult(3)
	unknown := judge.DiceResult(4)

	assert.Equal(t, "成功", success.String())
	assert.Equal(t, "失敗", fail.String())
	assert.Equal(t, "クリティカル", critical.String())
	assert.Equal(t, "ファンブル", fumble.String())
	assert.Equal(t, "異常な値です", unknown.String())
}

func TestValidateArgsSuccess(t *testing.T) {
	args := []string{"1", "2"}
	result := judgeCommand.ValidateArgs(args)

	assert.Equal(t, true, result)
}

func TestValidateArgsFail(t *testing.T) {
	args := []string{"hoge", "fuga"}
	result := judgeCommand.ValidateArgs(args)
	assert.Equal(t, false, result)

	args = []string{}

	result = judgeCommand.ValidateArgs(args)
	assert.Equal(t, false, result)
}

func TestGetPrefix(t *testing.T) {
	assert.Equal(t, "!judge", judgeCommand.GetPrefix())
}
