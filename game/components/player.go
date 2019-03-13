package components

import "fmt"

type Player struct {
	credits    int
	kill       int
	reputation int
	cargo      map[string]Goods
}

func (s *Player) Name() string {
	return "player"
}

func (s *Player) AddCargoGood(goods Goods) {
	// v是copy还是引用？
	v, ok := s.cargo[goods.name]
	if ok {
		v.quantity += goods.quantity
		s.cargo[goods.name] = v
	} else {
		s.cargo[goods.name] = goods
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
