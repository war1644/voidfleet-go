package systems

import (
	"fmt"
	"void_fleet/ecs"
	"void_fleet/game/components"
)

// Score ...
type Score struct{}

// NewScore ...
func NewScore() ecs.System {
	return &Score{}
}

// Process ...
func (s *Score) Update(world *ecs.World) {
	if world.Stop {
		return
	}
	scoreboard := world.GetEntity("scoreboard")
	score := scoreboard.GetComponent("score").(*components.Score)
	text := scoreboard.GetComponent("text").(*components.Text)
	if score.Enemy >= 10 {
		text.Content = "Enemy Wins!"
		score.Enemy = 0
		score.Player = 0
	} else if score.Player >= 10 {
		text.Content = "Player Wins!"
		score.Enemy = 0
		score.Player = 0
	} else {
		text.Content = fmt.Sprintf("%d : %d", score.Player, score.Enemy)
	}
}

// Setup ...
func (s *Score) Start() {}

// Teardown ...
func (s *Score) Remove() {}
