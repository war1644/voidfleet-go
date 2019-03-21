package game

type Galaxy struct {
	NameList    [][3]string
	List        map[string][]*Planet `json:"-"`
	Current     []*Planet
	CurrentName string
}

func NewGalaxy() *Galaxy {
	return &Galaxy{
		NameList:    [][3]string{},
		List:        make(map[string][]*Planet, 32),
		CurrentName: "天狼星区",
	}
}

func (s *Galaxy) GetRandPlanet(seed int) *Planet {
	return s.Current[RandNum(1, len(s.Current), seed)]
}

func (s *Galaxy) SetCurrent(name string) {
	s.Current = s.List[name]
}

//galaxy jump
func (s *Galaxy) Jump() {
	switch s.CurrentName {
	case "人马座":
		s.CurrentName = "烈阳星区"
	case "烈阳星区":
		s.CurrentName = "天狼星区"
	case "天狼星区":
		s.CurrentName = "北落师门"
	case "北落师门":
		s.CurrentName = "PLA"
	case "PLA":
		s.CurrentName = "X星区"
	case "X星区":
		s.CurrentName = "北极星区"
	case "北极星区":
		s.CurrentName = "人马座"
	default:
	}
}
