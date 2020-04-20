package game

import (
	"math/rand"
	"strconv"
	"time"
)

func RandNum(min, max, seed int) int {
	rand.Seed(time.Now().Unix() + int64(seed))
	randNum := rand.Intn(max-min) + min
	return randNum
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end-start) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func RandFloat(min, max int) float32 {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return float32(randNum)
}

func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func PlanetNamePool(seed int) ([2]string, int, int) {
	x := RandNum(0, 180, seed)
	y := RandNum(7, 164, seed)
	planetNumber := RandNum(1, 9, seed)
	/*星区图 0,7 80,7 0,64 80,64*/
	name := [][2]string{
		{"殖民星球" + strconv.Itoa(planetNumber), "star_1"},
		{"开发星球" + strconv.Itoa(planetNumber), "star_2"},
		{"空间站", "space_station"},
		{"海盗基地", "pirate_base"},
		{"海盗舰队", "pirate_fleet"},
		{"海盗小队", "pirate_group"},
		{"军事前哨", "outpost"},
		{"巡逻舰队", "outpost_fleet"},
		{"星系主力舰队", "main_fleet"},
	}
	i := RandNum(0, len(name), seed)
	return name[i], x, y
}
