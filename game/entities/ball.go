package entities

import (
	"ecs-pong/ecs"
	"ecs-pong/game/components"
)

// NewBall creates a new ball with an id on a specific position x and y with a custom width and height.
func NewBall(id string, x, y, width, height float32) (e *ecs.Entity) {
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
			&components.Sound{
				EventFilename: map[string]string{
					"collision": "assets/sounds/collision.wav",
				},
				Filename:  "",
				IsEnabled: true,
				Volume:    1.0,
			},
			&components.Texture{
				Filename:  "assets/textures/ball.png",
				IsEnabled: true,
			},
			&components.Velocity{
				X: -3,
				Y: 2,
			},
		},
	}
}
