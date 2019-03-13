package systems

import (
	"void_fleet/ecs"
	"void_fleet/game/components"
)

type AI struct{}

func NewAI() ecs.System {
	return &AI{}
}

// Process ...
func (s *AI) Update(world *ecs.World) {
	if world.Stop {
		return
	}
	ball := world.GetEntity("ball")
	for _, e := range world.FilterBy("ai", "position", "velocity") {
		s.handleBallPosition(e, ball)
	}
}

func (s *AI) Start() {}

func (s *AI) Remove() {}

func (s *AI) handleBallPosition(entity, ball *ecs.Entity) {
	ai := entity.GetComponent("ai").(*components.AI)
	position := entity.GetComponent("position").(*components.Position)
	velocity := entity.GetComponent("velocity").(*components.Velocity)
	ballPosition := ball.GetComponent("position").(*components.Position)
	if position.Y+velocity.Y < ballPosition.Y {
		ai.Down = true
		ai.Up = false
	}
	if position.Y+velocity.Y > ballPosition.Y {
		ai.Down = false
		ai.Up = true
	}
	if ai.Down {
		velocity.Y = 3
	}
	if ai.Up {
		velocity.Y = -3
	}
}
