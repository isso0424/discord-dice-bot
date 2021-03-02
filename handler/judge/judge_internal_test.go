package judge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const shouldOccurErrorMessage = "should occur error with %s"

func TestCompareResult(t *testing.T) {
	target := 50
	critical := 5
	success := 30
	fail := 70
	fumble := 96

	r := compareResult(target, critical)
	assert.Equal(t, "クリティカル", r.String())

	r = compareResult(target, success)
	assert.Equal(t, "成功", r.String())

	r = compareResult(target, fail)
	assert.Equal(t, "失敗", r.String())

	r = compareResult(target, fumble)
	assert.Equal(t, "ファンブル", r.String())
}

