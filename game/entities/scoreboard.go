package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

// NewScoreboard creates a new player with an id on a specific position x and y with a custom width and height.
func NewScoreboard(id string, x, y, width, height float32) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Size{
				Width:  width,
				Height: height,
			},
			&components.Text{
				Align:     components.TextAlignCenter,
				FontSize:  40,
				IsEnabled: true,
			},
			&components.Score{
				Enemy:  0,
				Player: 0,
			},
		},
	}
}
