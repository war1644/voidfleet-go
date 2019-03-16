package game

type Game struct {
	Galaxy        *Galaxy
	Player        *Player
	Items         *Items
	CurrentPlanet *Planet
	IsJump        bool
	Stop          bool
	Delay         int
	Load          *Load
}

func NewGame() *Game {
	game := &Game{
		Delay:  10,
		IsJump: true,
		Stop:   false,
		Galaxy: NewGalaxy(),
		Items:  NewItems(),
	}
	game.InitGalaxy()
	game.InitPlayer()
	return game
}

func (s *Game) InitPlayer() {
	s.Player = NewPlayer(999, s.Items.Ships[2], s.CurrentPlanet, s)
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
		planets[0] = NewPlanet("跳跃门", galaxyName, 40, 28)
		//设置跳跃门舰队（星系主力舰队）
		planets[0].Fleet = s.Items.SetRandomFleet(48)
		for i := 1; i < planetsLen; i++ {
			planetName, x, y := PlanetNamePool()
			xDistance := x - planets[0].X
			yDistance := y - planets[0].Y
			distance := xDistance + yDistance
			planets[i] = NewPlanet(planetName, galaxyName, x, y)
			//设置市场价
			planets[i].Goods = s.Items.SetRandomGoodsPrice()
			planets[i].Distance = distance
			//设置星球护卫舰队
			planets[i].Fleet = s.Items.SetRandomFleet(0)
		}
		s.Galaxy.List[galaxyName] = planets
	}
	s.CurrentPlanet = s.Galaxy.List[s.Galaxy.Current][0]

}
