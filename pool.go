package entitas

import "fmt"

type Pool interface {
	CreateEntity(cs ...Component) Entity
	Entities() []Entity
	Count() int
	HasEntity(e Entity) bool
	DestroyEntity(e Entity)
	DestroyAllEntities()
	Group(m Matcher) Group
}

type pool struct {
	index            int
	componentsLength ComponentType
	entities         map[EntityID]Entity
	groups           map[MatcherHash]Group
}

func NewPool(componentsLength ComponentType, index int) Pool {
	return &pool{
		index:            index,
		componentsLength: componentsLength,
		entities:         make(map[EntityID]Entity),
		groups:           make(map[MatcherHash]Group),
	}
}

func (p *pool) CreateEntity(cs ...Component) Entity {
	e := NewEntity(p.index)
	e.AddComponent(cs...)
	p.entities[e.ID()] = e
	p.index++
	e.AddCallback(ComponentAdded, p.componentAddedCallback)
	e.AddCallback(ComponentReplaced, p.componentReplacedCallback)
	e.AddCallback(ComponentRemoved, p.componentRemovedCallback)
	for _, g := range p.groups {
		g.HandleEntity(e)
	}
	return e
}

func (p *pool) Entities() []Entity {
	entities := make([]Entity, 0, len(p.entities))
	for _, e := range p.entities {
		entities = append(entities, e)
	}
	return entities
}

func (p *pool) Count() int {
	return len(p.entities)
}

func (p *pool) HasEntity(e Entity) bool {
	if entity, ok := p.entities[e.ID()]; ok && entity == e {
		return true
	}
	return false
}

func (p *pool) DestroyEntity(e Entity) {
	if entity, ok := p.entities[e.ID()]; ok && entity == e {
		e.RemoveAllComponents()
		delete(p.entities, e.ID())
		for _, g := range p.groups {
			g.HandleEntity(e)
		}
		return
	}
	panic("unknown entity")
}

func (p *pool) DestroyAllEntities() {
	for _, e := range p.entities {
		e.RemoveAllComponents()
	}
	p.entities = make(map[EntityID]Entity)
}

func (p *pool) Group(m Matcher) Group {
	if g, ok := p.groups[m.Hash()]; ok {
		return g
	}

	g := NewGroup(m)
	for _, e := range p.entities {
		g.HandleEntity(e)
	}
	p.groups[m.Hash()] = g
	return g
}

func (p *pool) String() string {
	return fmt.Sprintf("Pool(%v)", p.Entities())
}

func (p *pool) componentAddedCallback(e Entity, c Component) {
	for _, g := range p.groups {
		g.HandleEntity(e)
	}
}

func (p *pool) componentReplacedCallback(e Entity, c Component) {}

func (p *pool) componentRemovedCallback(e Entity, c Component) {
	for _, g := range p.groups {
		g.HandleEntity(e)
	}
}
