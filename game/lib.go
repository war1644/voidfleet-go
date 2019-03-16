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

func PlanetNamePool() (string, int, int) {
	x := RandNum(0, 80)
	y := RandNum(7, 64)
	planetNumber := RandNum(1, 500)
	/*星区图 0,7 80,7 0,64 80,64*/
	name := [9]string{
		"殖民星球" + string(planetNumber),
		"开发星球" + string(planetNumber),
		"空间站",
		"海盗基地",
		"海盗舰队",
		"海盗小队",
		"军事前哨",
		"巡逻舰队",
		"星系主力舰队",
	}
	i := RandNum(0, 8)
	return name[i], x, y
}
