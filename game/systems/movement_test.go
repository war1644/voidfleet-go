package systems_test

import (
	"ecs-pong/ecs"
	"ecs-pong/game/components"
	"ecs-pong/game/systems"
	"testing"
)

func TestMovement_Process_Position_Should_Not_Be_Changed_With_Velocity_Y_0(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 0},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_1_With_Velocity_Y_1(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 1},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
}

func TestMovement_Process_Position_Should_Be_Changed_To_Y_Minus_1_With_Velocity_Y_Minus_1(t *testing.T) {
	em := ecs.NewEntityManager()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: -1},
		},
	}
	em.Add(player)
	s := systems.NewMovement()
	s.Process(em)
}
