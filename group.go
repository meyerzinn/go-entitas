package entitas

type Group interface {
	Entities() []Entity
	Count() int
	HandleEntity(entity Entity)
}

type group struct {
	entities []Entity
}

func NewGroup(matcher Matcher) Group {
	return &group{}
}

func (g *group) Entities() []Entity {
	return g.entities
}

func (g *group) Count() int {
	return 0
}

func (g *group) HandleEntity(entity Entity) {
	g.entities = append(g.entities, entity)
}
