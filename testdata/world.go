package testdata

//var systems []System
//根管理
type World struct {
	systems  []System
	entities []*Entity
}

func New() *World {
	return &World{systems: make([]System, 64)}
}

func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
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
