package testdata

// game entity
var baseId uint64 = 0

type Component interface {
	Name() (name string)
}

type Entity struct {
	Components []Component
	id         uint64
	parent     *Entity
	children   []Entity
}

func (e *Entity) New() Entity {
	return Entity{id: baseId + 1}
}

func (e *Entity) Get(name string) Component {
	for _, c := range e.Components {
		if c.Name() == name {
			return c
		}
	}
	return nil
}

func (e *Entity) AppendChild(child *Entity) {
	child.parent = e
	e.children = append(e.children, *child)
}

func (e *Entity) Children() []Entity {
	return e.children
}

func (e *Entity) Parent() *Entity {
	return e.parent
}
