package data

import (
	"math/rand"
)

func getRandID() int {
	return int(rand.Int31n(999)) + 1
}
