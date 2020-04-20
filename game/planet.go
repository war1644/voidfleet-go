package game

type Planet struct {
	ID           string
	Name         string
	SpaceStation string
	Distance     int
	X            int
	Y            int
	Goods        []Goods `json:"-"`
	Fleet        []Ship  `json:"-"`
	EnemyFleet   []Ship
}

func NewPlanet(id, name string, x, y int) *Planet {
	return &Planet{
		ID:         id,
		Name:       name,
		X:          x,
		Y:          y,
		Fleet:      []Ship{},
		EnemyFleet: []Ship{},
		Goods:      []Goods{},
	}
}
