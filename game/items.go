package game

type Items struct {
	Goods []Goods
	Ships []Ship
}

func NewItems() *Items {
	my := &Items{}
	my.InitGameGoods()
	my.InitGameShip()
	return my
}

func (s *Items) InitGameShip() {
	s.Ships = []Ship{
		NewShip("帝国货船", "", 20000, 5000, 5, 25000, 200),
		NewShip("帝国采矿船", "", 15000, 1500, 5, 100000, 250),
		NewShip("帝国战机", "", 5000, 300, 10, 75000, 200),
		NewShip("帝国炮艇", "", 100000, 1000, 9, 125000, 250),
		NewShip("帝国护卫舰", "", 300000, 3000, 8, 525000, 300),
		NewShip("帝国驱逐舰", "", 1200000, 5000, 6, 2000000, 400),
		NewShip("帝国航母", "", 50100100, 10000, 5, 6000000, 500),
		NewShip("帝国歼星舰", "", 150100100, 20000, 2, 100000000, 1000),
	}
}

func (s *Items) InitGameGoods() {
	s.Goods = []Goods{
		//导弹
		NewGoods("台风", "", 200, true),
		NewGoods("黄蜂", "", 100, true),
		NewGoods("拳头", "", 300, true),
		NewGoods("鱼雷", "", 500, true),
		//装备
		NewGoods("跳跃引擎", "", 2000, true),
		NewGoods("主炮", "", 10000, true),
		NewGoods("机炮", "", 5000, true),
		NewGoods("激光炮", "", 20000, true),
		NewGoods("冲击波", "", 15000, true),
		//物品
		NewGoods("能量电池", "", 10, false),
		NewGoods("矿石", "", 30, false),
		NewGoods("金属", "", 50, false),
		NewGoods("武器", "", 100, false),
		NewGoods("营养液", "", 20, false),
		NewGoods("暗物质", "", 450000, false),
	}
}

//各星球价格随机，形成差价
func (s *Items) SetRandomGoodsPrice() []Goods {
	tmpGoods := make([]Goods, len(s.Goods))
	for i := range s.Goods {
		// copy value
		tmpGoods[i] = s.Goods[i]
		tmpGoods[i].Price = RandNum(tmpGoods[i].Price>>1, tmpGoods[i].Price<<1, i)
	}
	return tmpGoods
}

//为各星球生成舰队
func (s *Items) SetRandFleet(number int) []Ship {
	tmpNumber := 0
	if number > 0 {
		tmpNumber = number
	} else {
		tmpNumber = RandNum(12, 32, 1)
	}
	tmpShips := make([]Ship, tmpNumber)
	shipsLen := len(s.Ships)
	for i := 0; i < tmpNumber; i++ {
		tmpRand := RandNum(2, shipsLen, i)
		tmpShips[i] = s.Ships[tmpRand]
	}
	return tmpShips
}

//生成小队 战机、护卫舰
func (s *Items) SetRandGroup(number int) []Ship {
	tmpNumber := 0
	if number > 0 {
		tmpNumber = number
	} else {
		tmpNumber = RandNum(12, 32, 1)
	}
	tmpShips := make([]Ship, tmpNumber)
	shipsLen := len(s.Ships)
	for i := 0; i < tmpNumber; i++ {
		tmpRand := RandNum(2, shipsLen-3, i)
		tmpShips[i] = s.Ships[tmpRand]
	}
	return tmpShips
}
