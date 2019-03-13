package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewFleet(id string, x, y, width, height float32) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.AI{
				Up:    true,
				Down:  true,
				Left:  true,
				Right: true,
			},
			&components.Position{
				X: x, Y: y,
			},
			&components.Size{
				Width:  width,
				Height: height,
			},
			&components.Texture{
				Filename:  "assets/textures/paddle.png",
				IsEnabled: true,
			},
			&components.Velocity{
				Y:         0,
				X:         0,
				IsEnabled: true,
			},
		},
	}
}
