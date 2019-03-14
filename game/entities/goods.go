package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewGoods(id string, price int, describe string) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Goods{
				Price:     price,
				GoodsName: id,
				Describe:  describe,
			},
		},
	}
}
