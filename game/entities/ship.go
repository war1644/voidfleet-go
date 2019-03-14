package entities

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

func NewShip(id string, shipPrice int, shipCargo int, shipSpeed int, shipHp int, shipFuel int, shipDescribe string) (e *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Ship{
				ShipName: id,
				HP:       shipHp,
				MaxHp:    shipHp,
				Cargo:    shipCargo,
				Speed:    shipSpeed,
				Price:    shipPrice,
				Fuel:     shipFuel,
				MaxFuel:  shipFuel,
				Describe: shipDescribe,
			},
		},
	}
}
