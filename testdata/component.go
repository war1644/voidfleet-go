package testdata

import (
	"github.com/disintegration/gift"
	"image"
)

//type Component struct {
//	entity Entity
//}

type PositionComponent struct {
	Component
	x int
	y int
}
type SizeComponent struct {
	Component
	w int
	h int
}

type ColorComponent struct {
	Component
	r int
	g int
	b int
}

type SpeedComponent struct {
	Component
	speed float32
}

type HealthComponent struct {
	Component
	sp int
	ep int
}

type SpaceComponent struct {
	Component
	Width  float32
	Height float32
}

type CharacterComponent struct {
	name  string
	level int
}

type SpriteComponent struct {
	size     image.Rectangle // the sprite size
	Filter   *gift.GIFT      // normal filter used to draw the sprite
	FilterA  *gift.GIFT      // alternate filter used to draw the sprite
	FilterE  *gift.GIFT      // exploded filter used to draw the sprite
	Position image.Point     // top left position of the sprite
	Status   bool            // alive or dead
	Points   int             // number of points if destroyed
}
