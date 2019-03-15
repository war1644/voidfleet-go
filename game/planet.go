package game

type Planet struct {
	Name         string
	SpaceStation string
	Galaxy       string
	Distance     int
	X            int
	Y            int
	Goods        []Goods
	Fleet        []Ship
}

func NewPlanet(name, galaxy string, x, y int) *Planet {
	return &Planet{
		Name:   name,
		Galaxy: galaxy,
		X:      x,
		Y:      y,
		Fleet:  []Ship{},
		Goods:  []Goods{},
	}
}
