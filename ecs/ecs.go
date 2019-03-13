package ecs

type Component interface {
	Name() (name string)
}

type Entity struct {
	Components []Component
	Id         string
}

func (e *Entity) GetComponent(name string) Component {
	for _, c := range e.Components {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

type System interface {
	Start()
	Update(w *World)
	Remove()
}

type World struct {
	Stop     bool
	entities []*Entity
	systems  []System
}

func NewWorld() *World {
	return &World{
		Stop:     false,
		entities: []*Entity{},
		systems:  []System{},
	}
}

func (w *World) Update() {
	for _, system := range w.systems {
		system.Update(w)
	}
}

func (w *World) Start() {
	for _, system := range w.systems {
		system.Start()
	}
}

func (w *World) Remove() {
	for _, system := range w.systems {
		system.Remove()
	}
}

//manage Entity and System
func (w *World) AddEntity(entities ...*Entity) {
	for _, entity := range entities {
		w.entities = append(w.entities, entity)
	}
}

func (w *World) AddSystem(systems ...System) {
	for _, system := range systems {
		w.systems = append(w.systems, system)
	}
}

// 根据组件名 返回 实体
func (w *World) FilterBy(components ...string) (entities []*Entity) {
	for _, e := range w.entities {
		count := 0
		wanted := len(components)
		// Simply increase the count if the component could be found.
		for _, name := range components {
			for _, c := range e.Components {
				if c.Name() == name {
					count++
				}
			}
		}
		// Add the entity to the filter list, if all components are found.
		if count == wanted {
			entities = append(entities, e)
		}
	}
	return
}

func (w *World) GetEntity(id string) (entity *Entity) {
	for _, e := range w.entities {
		if e.Id == id {
			return e
		}
	}
	return
}

func (w *World) RemoveEntity(entity *Entity) {
	for i, e := range w.entities {
		if e.Id == entity.Id {
			copy(w.entities[i:], w.entities[i+1:])
			w.entities[len(w.entities)-1] = nil
			w.entities = w.entities[:len(w.entities)-1]
			break
		}
	}
}
