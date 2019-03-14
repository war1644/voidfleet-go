package systems

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

type Movement struct{}

func NewMovement() ecs.System {
	return &Movement{}
}

func (s *Movement) Update(world *ecs.World) {

	if world.Stop {
		return
	}
	for _, e := range world.FilterBy("position", "velocity") {
		position := e.GetComponent("position").(*components.Position)
		velocity := e.GetComponent("velocity").(*components.Velocity)
		position.X += velocity.X
		position.Y += velocity.Y
	}
}

func (s *Movement) Start(world *ecs.World) {}

func (s *Movement) Remove() {}
