package game

type Planet struct {
	Name         string
	SpaceStation string
	Galaxy       string
	Distance     int
	X            float32
	Y            float32
	Goods        []Goods
	Fleet        []Ship
}

func NewPlanet(name, galaxy string, x, y float32) *Planet {
	return &Planet{
		Name:   name,
		Galaxy: galaxy,
		X:      x,
		Y:      y,
	}
}
