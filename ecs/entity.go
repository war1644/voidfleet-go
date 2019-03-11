package ecs

// game entity
var (
	baseId uint64
)

func init() {
	baseId = 0
}

//type Identifier interface {
//	ID() uint64
//}
//
//func (e Entity) ID() uint64 {
//	return e.id
//}

type Entity struct {
	id       uint64
	parent   *Entity
	children []Entity
}

func (e *Entity) New() Entity {
	return Entity{id: baseId + 1}
}

func (e *Entity) GetEntity() *Entity {
	return e
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
