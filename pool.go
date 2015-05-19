package entitas

import (
	"container/list"
	"fmt"
)

type Pool interface {
	CreateEntity(components ...Component) Entity
	Entities() []Entity
	Count() int
	HasEntity(entity Entity) bool
	DestroyEntity(entity Entity)
	DestroyAllEntities()
	Group(matcher Matcher) Group
}

type pool struct {
	index            int
	componentsLength ComponentType
	entities         *list.List
}

func NewPool(componentsLength ComponentType, index int) Pool {
	return &pool{
		index:            index,
		entities:         list.New(),
		componentsLength: componentsLength,
	}
}

func (p *pool) CreateEntity(components ...Component) Entity {
	e := NewEntity(p.index)
	for _, c := range components {
		e.AddComponent(c)
	}
	p.entities.PushBack(e)
	p.index++
	return e
}

func (p *pool) Entities() []Entity {
	element := p.entities.Front()
	length := p.entities.Len()
	elements := make([]Entity, length)
	for i := 0; i < length; i++ {
		elements[i] = element.Value.(Entity)
		element = element.Next()
	}
	return elements
}

func (p *pool) Count() int {
	return len(p.Entities())
}

func (p *pool) HasEntity(entity Entity) bool {
	element := p.entities.Front()
	for {
		if element == nil {
			return false
		}
		if element.Value == entity {
			return true
		}
		element = element.Next()
	}
}

func (p *pool) DestroyEntity(entity Entity) {
	element := p.entities.Front()
	for {
		if element == nil {
			panic("tried to remove element not in list")
		}
		if element.Value == entity {
			p.entities.Remove(element)
			entity.RemoveAllComponents()
			return
		}
		element = element.Next()
	}
}

func (p *pool) DestroyAllEntities() {
	element := p.entities.Front()
	for element != nil {
		element.Value.(Entity).RemoveAllComponents()
		element = element.Next()
	}
	p.entities = p.entities.Init()
}

func (p *pool) Group(matcher Matcher) Group {
	g := NewGroup(matcher)
	element := p.entities.Front()
	for {
		if element == nil {
			break
		}
		g.HandleEntity(element.Value.(Entity))
		element = element.Next()
	}
	return g
}

func (p *pool) String() string {
	return fmt.Sprintf("Pool(%v)", p.Entities())
}
