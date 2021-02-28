package console_test

import (
	"fmt"
	"isso0424/dise/handler/console"
	"testing"
)

const shouldErrorOccur = "should error occur with %s"

func TestConsoleSend(t *testing.T) {
	session := console.New()
	err := session.Send("hoge", "fuga")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestConsoleSendFail(t *testing.T) {
	session := console.New()
	err := session.Send("", "message")
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "empty channel ID"))
	}

	err = session.Send("channelID", "")
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "empty message"))
	}

	err = session.Send("channelID", createTooLongMessage())
	if err == nil {
		t.Fatal(fmt.Sprintf(shouldErrorOccur, "too long message"))
	}
}

func createTooLongMessage() (message string) {
	template := "hoge"
	for i := 0; i < 510; i++ {
		message += template
	}

	return
}
