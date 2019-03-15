package systems_test

import (
	"fmt"
	"testing"
	"void_fleet/ecs"
	"void_fleet/game/components"
	"void_fleet/game/systems"
)

func TestMovement_a(t *testing.T) {
	world := ecs.NewWorld()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 0},
		},
	}
	s := systems.NewMovement()
	s.Update(world)
	fmt.Println(player.GetComponent("position"))
}

func TestMovement_System_Update(t *testing.T) {
	world := ecs.NewWorld()
	player := &ecs.Entity{
		Components: []ecs.Component{
			&components.Position{X: 0, Y: 0},
			&components.Velocity{Y: 0},
		},
	}
	s := systems.NewMovement()
	s.Update(world)
	fmt.Println(player.GetComponent("position"))
}

func TestMovement_System_Update_Change1(t *testing.T) {
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
	fmt.Println(player.GetComponent("position"))

}

func TestMovement_System_Update_Change2(t *testing.T) {
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
	fmt.Println(player.GetComponent("position"))
}
