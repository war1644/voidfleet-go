package testdata

import "image"

// game System
//type systems []System

type System interface {
	Update(dt float32)
	Remove(e Entity)
}

type MoveSystem struct {
}

func (s *MoveSystem) Update(dt float32) {

}

func (s *MoveSystem) Remove(e Entity) {

}

type CollideSystem struct {
}

func (s *CollideSystem) Update(dt float32) {

}

func (s *CollideSystem) Remove(e Entity) {

}

func (s *CollideSystem) check(s1, s2 SpriteComponent) bool {
	spriteA := image.Rect(s1.Position.X, s1.Position.Y, s1.Position.X+s1.size.Dx(), s1.Position.Y+s1.size.Dy())
	spriteB := image.Rect(s2.Position.X, s2.Position.Y, s2.Position.X+s1.size.Dx(), s2.Position.Y+s1.size.Dy())
	if spriteA.Min.X < spriteB.Max.X && spriteA.Max.X > spriteB.Min.X &&
		spriteA.Min.Y < spriteB.Max.Y && spriteA.Max.Y > spriteB.Min.Y {
		return true
	}
	return false
}
