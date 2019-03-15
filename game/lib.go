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

func RandFloat(min, max int) float32 {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return float32(randNum)
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
