package entitas

type Group interface {
	Entities() []Entity
	HandleEntity(entity Entity)
	ContainsEntity(entity Entity) bool
}

type group struct {
	entities []Entity
	matcher  Matcher
}

func NewGroup(matcher Matcher) Group {
	return &group{
		matcher: matcher,
	}
}

func (g *group) Entities() []Entity {
	return g.entities
}

func (g *group) HandleEntity(entity Entity) {
	if g.matcher.Matches(entity) {
		g.entities = append(g.entities, entity)
	} else {
		g.removeEntity(entity)
	}
}

func (g *group) ContainsEntity(entity Entity) bool {
	i := findEntity(g.entities, entity)
	if i == -1 {
		return false
	}
	return true
}

func (g *group) removeEntity(entity Entity) {
	i := findEntity(g.entities, entity)
	if i == -1 {
		panic("TODO")
	}
	copy(g.entities[i:], g.entities[i+1:])
	g.entities[len(g.entities)-1] = nil
	g.entities = g.entities[:len(g.entities)-1]
}

func findEntity(entities []Entity, entity Entity) int {
	for i, e := range entities {
		if e == entity {
			return i
		}
	}
	return -1
}
