package components

import "fmt"

type Player struct {
	Credits    int
	Kill       int
	Reputation int
	Cargo      map[string]Goods
	Year       int
	Day        int
	Planet     *Planet
	GoodsCount int
	Ship       Ship
}

func (s *Player) Name() string {
	return "player"
}

func (s *Player) AddCredits(c int) {
	s.Credits += c
}

func (s *Player) AddCargoGood(goods Goods) {
	// v是copy还是引用？
	v, ok := s.Cargo[goods.GoodsName]
	if ok {
		v.quantity += goods.quantity
		s.Cargo[goods.GoodsName] = v
	} else {
		s.Cargo[goods.GoodsName] = goods
	}
}

func (s *Player) GetState() {
	fmt.Println("--状态--",
		"Money：{credits}，旅行时间： {year} 年 {day} 天",
		"位置：{planet.galaxy} -- {planet.name}，旗舰：{ship.name}",
		"燃料： {ship.fuel}，飞行状态：{(!Game.docked).ToString()}",
		"装甲：{ship.hp}，杀敌：{kill}",
		"空余货仓：{(ship.cargo-goodsCount)}")
}
