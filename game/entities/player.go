package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewPlayer(id string, x, y, width, height float32) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Player{},
			&components.Size{
				Width:  width,
				Height: height,
			},
		},
	}
}
