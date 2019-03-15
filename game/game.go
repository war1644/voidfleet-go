package game

type Game struct {
	Goods  []Goods
	Ships  []Ship
	Galaxy *Galaxy
	IsJump bool
	Stop   bool
	Delay  int
	Load   *Load
}

func NewGame() *Game {
	game := &Game{
		Delay:  10,
		IsJump: true,
		Stop:   false,
		Goods:  []Goods{},
		Ships:  []Ship{},
		Galaxy: &Galaxy{
			NameList: []string{},
			List:     make(map[string][]*Planet, 32),
			Current:  "天狼星区",
		},
	}
	game.InitGameGoods()
	game.InitGameShip()
	game.InitGalaxy()
	return game
}

func (s *Game) InitGameShip() {
	s.Ships = []Ship{
		NewShip("帝国货船", "", 20000, 5000, 5, 25000, 200),
		NewShip("帝国采矿船", "", 15000, 1500, 5, 100000, 250),
		NewShip("帝国战斗机", "", 5000, 300, 10, 75000, 200),
		NewShip("帝国护卫舰", "", 300000, 3000, 8, 125000, 300),
		NewShip("帝国驱逐舰", "", 1200000, 5000, 6, 400000, 400),
		NewShip("帝国航母", "", 50100100, 10000, 5, 500000, 500),
		NewShip("帝国歼星舰", "", 150100100, 20000, 2, 1000000, 1000),
	}
}

func (s *Game) InitGameGoods() {
	s.Goods = []Goods{
		//导弹
		NewGoods("台风", "", 200, true),
		NewGoods("黄蜂", "", 100, true),
		NewGoods("拳头", "", 300, true),
		NewGoods("鱼雷", "", 500, true),
		//装备
		NewGoods("跳跃引擎", "", 20000, true),
		NewGoods("激光炮", "", 200000, true),
		NewGoods("主炮", "", 100000, true),
		NewGoods("冲击波炮", "", 150000, true),
		NewGoods("机炮", "", 50000, true),
		//物品
		NewGoods("能量电池", "", 10, false),
		NewGoods("金属", "", 50, false),
		NewGoods("矿石", "", 30, false),
		NewGoods("武器", "", 100, false),
		NewGoods("暗物质", "", 450000, false),
		NewGoods("兴奋剂", "", 1000, false),
	}
}

func (s *Game) InitGalaxy() {
	s.Galaxy.NameList = []string{
		"人马座",
		"烈阳星区",
		"天狼星区",
		"北落师门",
		"PLA",
		"北极星区",
	}
	for _, galaxyName := range s.Galaxy.NameList {
		planetsLen := RandNum(9, 18)
		planets := make([]*Planet, planetsLen)
		planets[0] = NewPlanet("星系跳跃门", galaxyName, 40, 28)
		//设置跳跃门舰队（星系主力舰队）
		planets[0].Fleet = s.SetRandomFleet(48)
		for i := 1; i < planetsLen; i++ {
			planetNumber := RandNum(1, 500)
			/*星区图 0,7 80,7 0,64 80,64*/
			x := RandNum(0, 80)
			y := RandNum(7, 64)
			xDistance := x - planets[0].X
			yDistance := y - planets[0].Y
			distance := xDistance + yDistance
			planets[i] = NewPlanet("开发行星"+string(planetNumber), galaxyName, x, y)
			//设置市场价
			planets[i].Goods = s.SetRandomGoodsPrice()[:]
			planets[i].Distance = distance
			//设置星球护卫舰队
			planets[i].Fleet = s.SetRandomFleet(0)
		}
		s.Galaxy.List[galaxyName] = planets
	}
}

//各星球价格随机，形成差价
func (s *Game) SetRandomGoodsPrice() []Goods {
	tmpGoods := make([]Goods, len(s.Goods))
	for i := range s.Goods {
		// copy value
		tmpGoods[i] = s.Goods[i]
		tmpGoods[i].Price = RandNum(tmpGoods[i].Price>>1, tmpGoods[i].Price<<1)
	}
	return tmpGoods
}

//为各星球生成舰队
func (s *Game) SetRandomFleet(number int) []Ship {
	tmpNumber := 0
	if number > 0 {
		tmpNumber = number
	} else {
		tmpNumber = RandNum(12, 32)
	}
	tmpShips := make([]Ship, tmpNumber)
	shipsLen := len(s.Ships)
	for i := 0; i < tmpNumber; i++ {
		tmpRand := RandNum(2, shipsLen)
		tmpShips[i] = s.Ships[tmpRand]
	}
	return tmpShips
}
