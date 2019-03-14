package systems

import (
	"void_fleet/ecs"
	"void_fleet/game/entities"
)

type Galaxy struct {
}

func (s *Galaxy) Start(world *ecs.World) {
	world.AddEntity(
		//导弹
		entities.NewGoods("台风", 200, ""),
		entities.NewGoods("黄蜂", 100, ""),
		entities.NewGoods("拳头", 300, ""),
		entities.NewGoods("鱼雷", 500, ""),
		//装备
		entities.NewGoods("跳跃引擎", 20000, ""),
		entities.NewGoods("激光炮", 200000, ""),
		entities.NewGoods("主炮", 100000, ""),
		entities.NewGoods("冲击波炮", 150000, ""),
		entities.NewGoods("拦截", 50000, ""),
		//物品
		entities.NewGoods("能量电池", 10, ""),
		entities.NewGoods("金属", 50, ""),
		entities.NewGoods("矿石", 30, ""),
		entities.NewGoods("武器", 100, ""),
		entities.NewGoods("暗物质", 450000, ""),
		entities.NewGoods("海洛因", 1000, ""),
	)

	world.AddEntity(
		entities.NewShip("帝国货船", 20000, 5000, 5, 25000, 200, ""),
		entities.NewShip("帝国采矿船", 15000, 1500, 5, 100000, 250, ""),
		entities.NewShip("帝国战斗机", 30000, 300, 10, 75000, 200, ""),
		entities.NewShip("帝国护卫舰", 300000, 3000, 8, 125000, 300, ""),
		entities.NewShip("帝国驱逐舰", 1200000, 5000, 6, 400000, 400, ""),
		entities.NewShip("帝国航母", 50100100, 10000, 5, 500000, 500, ""),
		entities.NewShip("帝国歼星舰", 150100100, 20000, 2, 1000000, 1000, ""),
	)
}

func (s *Galaxy) Update(world *ecs.World) {

}

func (s *Galaxy) Remove() {}
