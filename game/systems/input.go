package systems

import (
	"ecs-pong/ecs"
	"ecs-pong/game/components"
	"github.com/gen2brain/raylib-go/raylib"
)

// Input ...
type Input struct{}

// NewInput ...
func NewInput() ecs.System {
	return &Input{}
}

// Process ...
func (s *Input) Process(entityManager *ecs.EntityManager) {
	if rl.WindowShouldClose() {
		ecs.ShouldEngineStop = true
		return
	}
	for _, e := range entityManager.FilterBy("input", "velocity") {
		s.handleInput(e)
	}
}

// Setup ...
func (s *Input) Setup() {}

// Teardown ...
func (s *Input) Teardown() {}

func (s *Input) handleInput(e *ecs.Entity) {
	input := e.Get("input").(*components.Input)
	velocity := e.Get("velocity").(*components.Velocity)
	input.Down = rl.IsKeyDown(rl.KeyS)
	input.Up = rl.IsKeyDown(rl.KeyW)
	velocity.Y = 0
	if input.Down {
		velocity.Y = 4
	}
	if input.Up {
		velocity.Y = -4
	}
}
