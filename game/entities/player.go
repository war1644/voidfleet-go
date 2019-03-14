package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewPlayer(id string, x, y float32) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Position{
				X: x,
				Y: y,
			},
			&components.Player{
				Credits:    10000,
				Kill:       0,
				Reputation: 0,
				Cargo:      make(map[string]components.Goods),
				Year:       0,
				Day:        0,
				Planet:     &components.Planet{},
				GoodsCount: 0,
				Ship:       components.Ship{},
			},
		},
	}
}
