package dice

import (
	"math/rand"
	"time"
)

type Dice struct{}

func New() Dice {
	rand.Seed(time.Now().UnixNano())

	return Dice{}
}
