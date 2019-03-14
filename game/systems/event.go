package systems

import (
	"void_fleet/ecs"
)

type Event struct {
}

func NewEvent() ecs.System {
	return &Event{}
}

func (s *Event) Start(world *ecs.World) {
}

func (s *Event) Update(w *ecs.World) {

}

func (s *Event) Remove() {

}
