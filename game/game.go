package game

type Game struct {
	goods []Goods
	ships []Ship

}

var nameList = [6]string{
"人马座",
"烈阳星区",
"天狼星区",
"北落师门",
"PLA",
"北极星区",
}

func NewGame()  {

}

func (s *Game) Start() {

	s.goods = []Goods{
		//导弹
		NewGoods("台风","", 200),
		NewGoods("黄蜂", 100, ""),
		NewGoods("拳头", 300, ""),
		NewGoods("鱼雷", 500, ""),
		//装备
		NewGoods("跳跃引擎", 20000, ""),
		NewGoods("激光炮", 200000, ""),
		NewGoods("主炮", 100000, ""),
		NewGoods("冲击波炮", 150000, ""),
		NewGoods("拦截", 50000, ""),
		//物品
		NewGoods("能量电池", 10, ""),
		NewGoods("金属", 50, ""),
		NewGoods("矿石", 30, ""),
		NewGoods("武器", 100, ""),
		NewGoods("暗物质", 450000, ""),
		NewGoods("海洛因", 1000, ""),
	}

	ships := []Ship{
		NewShip("帝国货船", 20000, 5000, 5, 25000, 200, ""),
		NewShip("帝国采矿船", 15000, 1500, 5, 100000, 250, ""),
		NewShip("帝国战斗机", 30000, 300, 10, 75000, 200, ""),
		NewShip("帝国护卫舰", 300000, 3000, 8, 125000, 300, ""),
		NewShip("帝国驱逐舰", 1200000, 5000, 6, 400000, 400, ""),
		NewShip("帝国航母", 50100100, 10000, 5, 500000, 500, ""),
		NewShip("帝国歼星舰", 150100100, 20000, 2, 1000000, 1000, ""),
	}
}

func NewGalaxy() {
	for _,galaxyName := range nameList {
		planets := [RandNum(5,9)]*Planet{}
		planets[0] = NewPlanet("星系跳跃门",galaxyName, 0, 0)

		//设置跳跃门舰队（星系主力舰队）
		List<Ship> masterFleet1 = SetRandomFleet()
		List<Ship> masterFleet2 = SetRandomFleet()
		masterFleet1.AddRange(masterFleet2)
		planet[0].fleet = masterFleet1

		for  i := 1 i < planet.Length i++ {
			int planetNumber = rand.Next(1, 500)
			int x = rand.Next(-50, 50)
			int y = rand.Next(-50, 50)
			int xDistance = x - planet[0].x
			int yDistance = y - planet[0].y
			int distance = xDistance + yDistance
			planet[i] = new Planet($"开发行星{planetNumber}",galaxyName,x,y)
			//设置市场价
			Good[] localGoods = SetRandomGoodsPrice()
			planet[i].goods = localGoods
			planet[i].distance = distance

			//设置星球护卫舰队
			List<Ship> localFleet = SetRandomFleet()
			planet[i].fleet = localFleet
		}
	}
	s.List.Add(galaxyName,planet)
}

//各星球价格随机，形成差价
func SetRandomGoodsPrice() []Goods {
	tmpGoods := NewGoods[goods.Length]
	for (int i = 0 i < goods.Length i++){
		tmpGoods[i] = Clone<Good>(goods[i])
		tmpGoods[i].price = rand.Next(tmpGoods[i].price/10, tmpGoods[i].price)
	}
	return tmpGoods
}

//为各星球生成舰队
func (s *Game) SetRandomFleet() [...]Ship {
	tmpNumber := RandNum(1,9)
	tmpShips := [tmpNumber]Ship{}
	for i := 0; i < tmpNumber; i++ {
		tmpRand := RandNum(2, len(s.ships))
		copy(s.ships[tmpRand],tmpShips)
	}
	return tmpShips
}
