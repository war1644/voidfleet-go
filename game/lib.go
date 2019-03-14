package game

import (
	"math/rand"
	"time"
)

func RandNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
