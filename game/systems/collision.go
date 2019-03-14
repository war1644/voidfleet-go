package systems

import (
	"image"
	"void_fleet/ecs"
	"void_fleet/game/components"
)

// Collision ...
type Collision struct {
	windowHeight int32
	windowWidth  int32
}

// NewCollision ...
func NewCollision(windowWidth, windowHeight int32) ecs.System {
	return &Collision{
		windowHeight: windowHeight,
		windowWidth:  windowWidth,
	}
}

// Process ...
func (s *Collision) Update(world *ecs.World) {

	if world.Stop {
		return
	}
	for _, e := range world.FilterBy("position", "size", "velocity") {
		switch e.Id {
		case "ball":
			enemy := world.GetEntity("enemy")
			player := world.GetEntity("player")
			scoreboard := world.GetEntity("scoreboard")
			if s.hasCollisionWithEnemy(e, enemy) ||
				s.hasCollisionWithPlayer(e, player) ||
				s.hasCollisionWithWindowBottom(e) ||
				s.hasCollisionWithWindowTop(e) {
				s.handleCollisionSoundIfPresent(e)
			}
			s.handleEnemyScore(e, enemy, scoreboard)
			s.handlePlayerScore(e, player, scoreboard)
		case "enemy", "player":
			s.blockWindowBottom(e)
			s.blockWindowTop(e)
		}
	}
}

// Setup ...
func (s *Collision) Start(world *ecs.World) {}

// Teardown ...
func (s *Collision) Remove() {}

func (s *Collision) blockWindowBottom(entity *ecs.Entity) {
	position := entity.GetComponent("position").(*components.Position)
	size := entity.GetComponent("size").(*components.Size)
	velocity := entity.GetComponent("velocity").(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y = 0
	}
}

func (s *Collision) blockWindowTop(entity *ecs.Entity) {
	position := entity.GetComponent("position").(*components.Position)
	velocity := entity.GetComponent("velocity").(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y = 0
	}
}

func (s *Collision) getEntityRect(entity *ecs.Entity) components.Rect {
	position := entity.GetComponent("position").(*components.Position)
	size := entity.GetComponent("size").(*components.Size)
	return components.Rect{X: int(position.X), Y: int(position.Y), W: int(size.Width), H: int(size.Height)}
}

func (s *Collision) handleCollisionSoundIfPresent(ball *ecs.Entity) {
	sound := ball.GetComponent("sound")
	if sound == nil {
		return
	}
	snd := sound.(*components.Sound)
	snd.Filename = snd.EventFilename["collision"]
}

func (s *Collision) hasCollisionWithEnemy(ball, enemy *ecs.Entity) (hasCollision bool) {
	ballRect := s.getEntityRect(ball)
	ballVelocity := ball.GetComponent("velocity").(*components.Velocity)
	enemyRect := s.getEntityRect(enemy)
	enemyAI := enemy.GetComponent("ai").(*components.AI)
	if s.check(ballRect, enemyRect) {
		ballVelocity.X *= -1
		if enemyAI.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 2
		} else if enemyAI.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= 0.5
			ballVelocity.X *= 1.5
		} else if enemyAI.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 2
		} else if enemyAI.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= 0.5
			ballVelocity.X *= 1.5
		}
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithPlayer(ball, player *ecs.Entity) (hasCollision bool) {
	ballRect := s.getEntityRect(ball)
	ballVelocity := ball.GetComponent("velocity").(*components.Velocity)
	playerRect := s.getEntityRect(player)
	playerInput := player.GetComponent("input").(*components.Input)
	if s.check(ballRect, playerRect) {
		ballVelocity.X *= -1
		if playerInput.Down && ballVelocity.Y > 0 {
			ballVelocity.Y *= 2
		} else if playerInput.Down && ballVelocity.Y < 0 {
			ballVelocity.Y *= -0.5
			ballVelocity.X *= 1.5
		} else if playerInput.Up && ballVelocity.Y < 0 {
			ballVelocity.Y *= 2
		} else if playerInput.Up && ballVelocity.Y > 0 {
			ballVelocity.Y *= -0.5
			ballVelocity.X *= 1.5
		}
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithWindowBottom(entity *ecs.Entity) (hasCollision bool) {
	position := entity.GetComponent("position").(*components.Position)
	size := entity.GetComponent("size").(*components.Size)
	velocity := entity.GetComponent("velocity").(*components.Velocity)
	if position.Y+velocity.Y+size.Height >= float32(s.windowHeight) {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *Collision) hasCollisionWithWindowTop(entity *ecs.Entity) (hasCollision bool) {
	position := entity.GetComponent("position").(*components.Position)
	velocity := entity.GetComponent("velocity").(*components.Velocity)
	if position.Y+velocity.Y <= 0 {
		velocity.Y *= -1
		return true
	}
	return false
}

func (s *Collision) handleEnemyScore(ball, enemy, scoreboard *ecs.Entity) {
	position := ball.GetComponent("position").(*components.Position)
	velocity := ball.GetComponent("velocity").(*components.Velocity)
	score := scoreboard.GetComponent("score").(*components.Score)
	if position.X+velocity.X <= 0 {
		score.Enemy++
		velocity.X = -3
		velocity.Y = 2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}

func (s *Collision) handlePlayerScore(ball, player, scoreboard *ecs.Entity) {
	position := ball.GetComponent("position").(*components.Position)
	velocity := ball.GetComponent("velocity").(*components.Velocity)
	score := scoreboard.GetComponent("score").(*components.Score)
	if position.X+velocity.X >= float32(s.windowWidth) {
		score.Player++
		velocity.X = 3
		velocity.Y = -2
		position.X = float32(s.windowWidth) / 2
		position.Y = float32(s.windowHeight) / 2
	}
}

func (s *Collision) check(s1, s2 components.Rect) bool {
	spriteA := image.Rect(s1.X, s1.Y, s1.X+s1.W, s1.Y+s1.H)
	spriteB := image.Rect(s2.X, s2.Y, s2.X+s1.W, s2.Y+s1.H)
	if spriteA.Min.X < spriteB.Max.X && spriteA.Max.X > spriteB.Min.X &&
		spriteA.Min.Y < spriteB.Max.Y && spriteA.Max.Y > spriteB.Min.Y {
		return true
	}
	return false
}
