package systems_test

import (
	"testing"
	"void_fleet/ecs"
	"void_fleet/game/components"
	"void_fleet/game/systems"
)

func TestMovement_Process_Position_Should_Not_Be_Changed_With_Velocity_Y_0(t *testing.T) {
	world := ecs.NewWorld()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 0},
		},
	}
	world.AddEntity(player)
	s := systems.NewMovement()
	s.Update(world)
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_1_With_Velocity_Y_1(t *testing.T) {
	world := ecs.NewWorld()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 1},
		},
	}
	world.AddEntity(player)
	s := systems.NewMovement()
	s.Update(world)
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_Minus_1_With_Velocity_Y_Minus_1(t *testing.T) {
	world := ecs.NewWorld()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: -1},
		},
	}
	world.AddEntity(player)
	s := systems.NewMovement()
	s.Update(world)
}
