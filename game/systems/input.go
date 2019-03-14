package systems

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

// Input ...
type Input struct{}

// NewInput ...
func NewInput() ecs.System {
	return &Input{}
}

// Process ...
func (s *Input) Update(world *ecs.World) {

	for _, e := range world.FilterBy("input", "velocity") {
		s.handleInput(e)
	}
}

// Setup ...
func (s *Input) Start(world *ecs.World) {}

// Teardown ...
func (s *Input) Remove() {}

func (s *Input) handleInput(e *ecs.Entity) {
	input := e.GetComponent("input").(*components.Input)
	velocity := e.GetComponent("velocity").(*components.Velocity)
	input.Down = false
	input.Up = true
	velocity.Y = 0
	if input.Down {
		velocity.Y = 4
	}
	if input.Up {
		velocity.Y = -4
	}
}
