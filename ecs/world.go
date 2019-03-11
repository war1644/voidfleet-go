package ecs

import (
	"reflect"
)

type World struct {
	systems      systems
	sysIn, sysEx map[reflect.Type]reflect.Type
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
	//sort.Sort(w.systems)
}

func (w *World) Systems() []System {
	return w.systems
}

func (w *World) Update(dt float32) {
	for _, system := range w.Systems() {
		system.Update(dt)
	}
}

func (w *World) RemoveEntity(e Entity) {
	for _, sys := range w.systems {
		sys.Remove(e)
	}
}
