package game

type Game struct {
	Galaxy        *Galaxy
	Player        *Player
	Items         *Items
	Event         *Event
	CurrentPlanet *Planet
	IsJump        bool
	Stop          bool
	Delay         int
	Load          *Load
	MsgType       [3]string
}

func NewGame() *Game {
	game := &Game{
		//info 行为信息 primary 新闻 success任务
		MsgType: [3]string{"info", "primary", "success"},
		Delay:   10,
		IsJump:  true,
		Stop:    false,
		Galaxy:  NewGalaxy(),
		Items:   NewItems(),
		Event:   NewEvent(),
	}
	game.InitGalaxy()
	game.InitPlayer()
	//生成一些随机事件
	game.randEvent()
	return game
}

func (s *Game) InitPlayer() {
	s.Player = NewPlayer(999, s.Items.Ships[2], s.CurrentPlanet, s)
	//加些物品
	goods1 := s.Items.Goods[5-1]
	goods2 := s.Items.Goods[7-1]
	goods3 := s.Items.Goods[11-1]
	goods1.Quantity = 1
	goods2.Quantity = 2
	goods3.Quantity = 99
	s.Player.AddCargoGoods(goods1, goods2, goods3)
}

func (s *Game) InitGalaxy() {

	s.Galaxy.NameList = [][3]string{
		{"人马座", "75", "50"},
		{"烈阳星区", "30", "33"},
		{"天狼星区", "50", "10"},
		{"北落师门", "7", "18"},
		{"PLA星区", "40", "25"},
		{"X星区", "5", "45"},
		{"北极星区", "63", "22"},
	}
	for i, galaxyName := range s.Galaxy.NameList {
		planetsLen := RandNum(9, 18, i)
		planets := make([]*Planet, planetsLen)
		planets[0] = NewPlanet("gate", "跳跃门", 40, 28)
		//设置跳跃门舰队（星系主力舰队）
		planets[0].Fleet = s.Items.SetRandFleet(48)
		for i := 1; i < planetsLen; i++ {
			idName, x, y := PlanetNamePool(i)
			xDistance := x - planets[0].X
			yDistance := y - planets[0].Y
			distance := xDistance + yDistance
			planets[i] = NewPlanet(idName[1], idName[0], x, y)
			//设置市场价
			planets[i].Goods = s.Items.SetRandomGoodsPrice()
			planets[i].Distance = distance
			//设置星球护卫舰队
			planets[i].Fleet = s.Items.SetRandFleet(0)
		}
		s.Galaxy.List[galaxyName[0]] = planets
	}
	s.Galaxy.CurrentName = "天狼星区"
	s.CurrentPlanet = s.Galaxy.List[s.Galaxy.CurrentName][0]
	s.Galaxy.SetCurrent(s.Galaxy.CurrentName)

	//星区星系生成完成，输出信息
	s.Event.NewMsg(s.MsgType[0], "到达:"+s.Galaxy.CurrentName+"-"+s.CurrentPlanet.Name)
}

func (s *Game) randEvent() {
	//海盗袭击
	p1 := s.Galaxy.GetRandPlanet(1)
	s.Event.NewMsg(s.MsgType[1], "海盗正在袭击:"+s.Galaxy.CurrentName+"-"+p1.Name)
	p1.EnemyFleet = s.Items.SetRandGroup(6)
	s.Event.NewMsg(s.MsgType[1], s.Galaxy.CurrentName+"-"+p1.Name+"护卫舰队正在迎战")

	//势力战争
	//p2 := s.Galaxy.GetRandPlanet(2)
	//s.Event.NewMsg(s.MsgType[1], "海盗正在袭击:"+s.Galaxy.CurrentName+"-"+p2.Name)
	//p2.EnemyFleet = s.Items.SetRandGroup(6)
	//s.Event.NewMsg(s.MsgType[1], s.Galaxy.CurrentName+"-"+p2.Name+"护卫舰队正在迎战")

	//贸易
	p3 := s.Galaxy.GetRandPlanet(3)
	s.Event.NewMsg(s.MsgType[1], s.Galaxy.CurrentName+"-"+p3.Name+"急缺"+p3.Goods[10].Name)
	p4 := s.Galaxy.GetRandPlanet(4)
	s.Event.NewMsg(s.MsgType[1], s.Galaxy.CurrentName+"-"+p4.Name+p4.Goods[12].Name+"滞销")

}
