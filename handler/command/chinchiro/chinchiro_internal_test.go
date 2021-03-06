package chinchiro

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJudge(t *testing.T) {
	assert.Equal(t, chinchiroResult(pinZoro), judge([]int{1, 1, 1}))

	assert.Equal(t, chinchiroResult(demeNi), judge([]int{1, 2, 1}))

	assert.Equal(t, chinchiroResult(hihumi), judge([]int{3, 2, 1}))

	assert.Equal(t, chinchiroResult(shigoro), judge([]int{5, 4, 6}))

	assert.Equal(t, chinchiroResult(yakuNashi), judge([]int{1, 2, 4}))
}

func TestZoro(t *testing.T) {
	assert.Equal(t, chinchiroResult(pinZoro), zoro(1))
	assert.Equal(t, chinchiroResult(niZoro), zoro(2))
	assert.Equal(t, chinchiroResult(sanZoro), zoro(3))
	assert.Equal(t, chinchiroResult(yonZoro), zoro(4))
	assert.Equal(t, chinchiroResult(goZoro), zoro(5))
	assert.Equal(t, chinchiroResult(rokuZoro), zoro(6))
	assert.Equal(t, chinchiroResult(yakuNashi), zoro(7))
}

func TestDeme(t *testing.T) {
	assert.Equal(t, chinchiroResult(demeIchi), deme([]int{2, 2, 1}))
	assert.Equal(t, chinchiroResult(demeNi), deme([]int{1, 2, 1}))
	assert.Equal(t, chinchiroResult(demeSan), deme([]int{1, 1, 3}))
	assert.Equal(t, chinchiroResult(demeYon), deme([]int{4, 1, 1}))
	assert.Equal(t, chinchiroResult(demeGo), deme([]int{5, 1, 1}))
	assert.Equal(t, chinchiroResult(demeRoku), deme([]int{6, 1, 1}))
	assert.Equal(t, chinchiroResult(yakuNashi), deme([]int{7, 1, 1}))
}

func TestChinchiroResult(t *testing.T) {
	assert.Equal(t, "ピンゾロ", chinchiroResult(pinZoro).String())
	assert.Equal(t, "ニゾロ", chinchiroResult(niZoro).String())
	assert.Equal(t, "サンゾロ", chinchiroResult(sanZoro).String())
	assert.Equal(t, "ヨンゾロ", chinchiroResult(yonZoro).String())
	assert.Equal(t, "ゴゾロ", chinchiroResult(goZoro).String())
	assert.Equal(t, "ロクゾロ", chinchiroResult(rokuZoro).String())
	assert.Equal(t, "ヒフミ", chinchiroResult(hihumi).String())
	assert.Equal(t, "シゴロ", chinchiroResult(shigoro).String())
	assert.Equal(t, "出目イチ", chinchiroResult(demeIchi).String())
	assert.Equal(t, "出目ニ", chinchiroResult(demeNi).String())
	assert.Equal(t, "出目サン", chinchiroResult(demeSan).String())
	assert.Equal(t, "出目ヨン", chinchiroResult(demeYon).String())
	assert.Equal(t, "出目ゴ", chinchiroResult(demeGo).String())
	assert.Equal(t, "出目ロク", chinchiroResult(demeRoku).String())
	assert.Equal(t, "役無し", chinchiroResult(yakuNashi).String())
	assert.Equal(t, "エラー", chinchiroResult(-1).String())
}
