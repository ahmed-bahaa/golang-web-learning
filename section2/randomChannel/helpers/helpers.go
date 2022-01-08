package helpers

import (
	"math/rand"
	"time"
)

func GetRandomInt(rng int) int {

	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(rng)
	return value

}
