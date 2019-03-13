package components

type Planet struct {
	name         string
	SpaceStation string
	galaxy       string
	distance     int
}

// Name ...
func (i *Planet) Name() string {
	return "planet"
}
