package entitas

import (
	"errors"
	"fmt"
	"sort"
)

var (
	ErrComponentExists       = errors.New("component exists")
	ErrComponentDoesNotExist = errors.New("component does not exist")
)

type EntityID uint

type Entity interface {
	ID() EntityID
	AddComponent(cs ...Component) error
	HasComponent(ts ...ComponentType) bool
	HasAnyComponent(ts ...ComponentType) bool
	RemoveComponent(ts ...ComponentType)
	Component(t ComponentType) (Component, error)
	ReplaceComponent(cs ...Component)
	Components() []Component
	ComponentIndices() []ComponentType
	RemoveAllComponents()
}

type entity struct {
	id         EntityID
	components map[ComponentType]Component
}

func NewEntity(id int) Entity {
	return &entity{
		id:         EntityID(id),
		components: make(map[ComponentType]Component),
	}
}

func (e *entity) ID() EntityID {
	return e.id
}

func (e *entity) AddComponent(cs ...Component) error {
	for _, c := range cs {
		if e.HasComponent(c.Type()) {
			return ErrComponentExists
		}
		e.components[c.Type()] = c
	}
	return nil
}

func (e *entity) HasComponent(ts ...ComponentType) bool {
	for _, t := range ts {
		if _, ok := e.components[t]; !ok {
			return false
		}
	}
	return true
}

func (e *entity) HasAnyComponent(ts ...ComponentType) bool {
	for _, t := range ts {
		if _, ok := e.components[t]; ok {
			return true
		}
	}
	return false
}

func (e *entity) RemoveComponent(ts ...ComponentType) {
	for _, t := range ts {
		delete(e.components, t)
	}
}

func (e *entity) Component(t ComponentType) (Component, error) {
	c := e.components[t]
	if c == nil {
		return nil, ErrComponentDoesNotExist
	}
	return c, nil
}

func (e *entity) ReplaceComponent(cs ...Component) {
	for _, c := range cs {
		e.components[c.Type()] = c
	}
}

func (e *entity) Components() []Component {
	components := make([]Component, len(e.components))
	i := 0
	for _, c := range e.components {
		components[i] = c
		i++
	}
	sort.Sort(ComponentsByType(components))
	return components
}

func (e *entity) ComponentIndices() []ComponentType {
	types := make([]ComponentType, len(e.components))
	i := 0
	for t := range e.components {
		types[i] = t
		i++
	}
	return types
}

func (e *entity) RemoveAllComponents() {
	e.components = make(map[ComponentType]Component)
}

func (e *entity) String() string {
	return fmt.Sprintf("Entity_%d(%v)", e.id, e.Components())
}
