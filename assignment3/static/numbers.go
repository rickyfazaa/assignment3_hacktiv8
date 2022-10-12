package static

import (
	"math/rand"
)

func RandomNumberWater() int {
	number := rand.Intn(99) + 1
	return number
}

func RandomNumberWind() int {
	number := rand.Intn(99) + 1
	return number
}
