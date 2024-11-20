package property

import (
	"math/rand"
	"time"
)

func GeneratePropertyUniqueId() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000000000
	max := 9999999999999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}
