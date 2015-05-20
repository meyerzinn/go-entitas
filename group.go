package entitas

type Group interface {
	Entities() []Entity
	HandleEntity(e Entity)
	ContainsEntity(e Entity) bool
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

func (g *group) HandleEntity(e Entity) {
	i := findEntity(g.entities, e)
	if i == -1 {
		if g.matcher.Matches(e) {
			g.entities = append(g.entities, e)
		}
	} else {
		if !g.matcher.Matches(e) {
			g.removeEntity(i)
		}
	}
}

func (g *group) ContainsEntity(e Entity) bool {
	if findEntity(g.entities, e) == -1 {
		return false
	}
	return true
}

func findEntity(entities []Entity, e Entity) int {
	for i, entity := range entities {
		if entity == e {
			return i
		}
	}
	return -1
}

func (g *group) removeEntity(i int) {
	copy(g.entities[i:], g.entities[i+1:])
	g.entities[len(g.entities)-1] = nil
	g.entities = g.entities[:len(g.entities)-1]
}
