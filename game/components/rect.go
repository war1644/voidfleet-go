package components

// Position contains the 2D X and Y coordinate.
type Rect struct {
	X int
	Y int
	W int
	H int
}

// Name ...
func (p *Rect) Name() string {
	return "rect"
}
