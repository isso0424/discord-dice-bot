package dice

import (
	"time"
	"math/rand"
)

type Dice struct { }

func New() Dice {
	rand.Seed(time.Now().UnixNano())

	return Dice {}
}
