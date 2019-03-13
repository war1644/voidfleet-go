package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewGalay(id string, x, y, width, height float32) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Galaxy{
				NameList: [6]string{"人马座", "烈阳星区", "天狼星区", "北落师门", "PLA", "北极星区"},
				List:     make(map[string]*components.Planet),
				Current:  "天狼星区",
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
				Y: 0,
			},
		},
	}
}
