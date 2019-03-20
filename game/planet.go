package game

type Planet struct {
	ID           string
	Name         string
	SpaceStation string
	Galaxy       string
	Distance     int
	X            int
	Y            int
	Goods        []Goods `json:"-"`
	Fleet        []Ship  `json:"-"`
}

func NewPlanet(id, name, galaxy string, x, y int) *Planet {
	return &Planet{
		ID:     id,
		Name:   name,
		Galaxy: galaxy,
		X:      x,
		Y:      y,
		Fleet:  []Ship{},
		Goods:  []Goods{},
	}
}
