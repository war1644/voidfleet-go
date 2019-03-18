package game

type Galaxy struct {
	NameList []string
	List     map[string][]*Planet `json:"-"`
	Current  string
}

func NewGalaxy() *Galaxy {
	return &Galaxy{
		NameList: []string{},
		List:     make(map[string][]*Planet, 32),
		Current:  "天狼星区",
	}
}

//galaxy jump
func (s *Galaxy) Jump() {
	switch s.Current {
	case "人马座":
		s.Current = "烈阳星区"
	case "烈阳星区":
		s.Current = "天狼星区"
	case "天狼星区":
		s.Current = "北落师门"
	case "北落师门":
		s.Current = "PLA"
	case "PLA":
		s.Current = "北极星区"
	case "北极星区":
		s.Current = "人马座"
	default:
	}
}
