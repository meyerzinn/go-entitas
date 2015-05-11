package entitas

import (
	"errors"
	"fmt"
	"sort"
)

type Entity interface {
	Index() int
	AddComponent(component ...Component)
	HasComponent(types ...ComponentType) bool
	HasAnyComponent(types ...ComponentType) bool
	RemoveComponent(types ...ComponentType)
	Component(t ComponentType) (Component, error)
	ReplaceComponent(components ...Component)
	Components() []Component
	ComponentIndices() []ComponentType
	RemoveAllComponents()
}

type entity struct {
	index      int
	components map[ComponentType]Component
}

func NewEntity(index int, componentsLength ComponentType) Entity {
	return &entity{
		index:      index,
		components: make(map[ComponentType]Component),
	}
}

func (e *entity) Index() int {
	return e.index
}

func (e *entity) AddComponent(components ...Component) {
	for _, c := range components {
		e.components[c.Type()] = c
	}
}

func (e *entity) HasComponent(types ...ComponentType) bool {
	for _, t := range types {
		if e.components[t] == nil {
			return false
		}
	}
	return true
}

func (e *entity) HasAnyComponent(types ...ComponentType) bool {
	for _, t := range types {
		if e.components[t] != nil {
			return true
		}
	}
	return false
}

func (e *entity) RemoveComponent(types ...ComponentType) {
	for _, t := range types {
		e.components[t] = nil
	}
}

func (e *entity) Component(t ComponentType) (Component, error) {
	c := e.components[t]
	if c == nil {
		return nil, errors.New("component not found")
	}
	return c, nil
}

func (e *entity) ReplaceComponent(components ...Component) {
	for _, c := range components {
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
	return fmt.Sprintf("Entity_%d(%v)", e.index, e.Components())
}
