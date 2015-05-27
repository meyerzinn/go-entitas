package entitas

type Group interface {
	Entities() []Entity
	HandleEntity(e Entity)
	ContainsEntity(e Entity) bool
}

type group struct {
	entities map[EntityID]Entity
	cache    []Entity
	matcher  Matcher
}

func NewGroup(matcher Matcher) Group {
	return &group{
		entities: make(map[EntityID]Entity),
		cache:    make([]Entity, 0),
		matcher:  matcher,
	}
}

func (g *group) Entities() []Entity {
	return g.cache
}

func (g *group) HandleEntity(e Entity) {
	if g.matcher.Matches(e) {
		g.addEntity(e)
	} else {
		g.removeEntity(e)
	}
}

func (g *group) addEntity(e Entity) {
	g.entities[e.ID()] = e
	g.cache = append(g.cache, e)
}

func (g *group) removeEntity(e Entity) {
	delete(g.entities, e.ID())
	if i := findIndex(g.cache, e); i != -1 {
		g.cache = removeIndexed(g.cache, i)
	}
}

func (g *group) ContainsEntity(e Entity) bool {
	if _, ok := g.entities[e.ID()]; ok {
		return true
	}
	return false
}

func findIndex(entities []Entity, e Entity) int {
	for i, entity := range entities {
		if entity == e {
			return i
		}
	}
	return -1
}

func removeIndexed(entities []Entity, i int) []Entity {
	copy(entities[i:], entities[i+1:])
	entities[len(entities)-1] = nil
	return entities[:len(entities)-1]
}
