package entitas

type Group interface {
	Entities() []Entity
	HandleEntity(e Entity)
	ContainsEntity(e Entity) bool
	AddCallback(e GroupEvent, c GroupCallback)
}

type GroupEvent uint

const (
	EntityAdded GroupEvent = iota
	EntityRemoved
)

type GroupCallback func(Group, Entity)

type group struct {
	entities  map[EntityID]Entity
	cache     []Entity
	matcher   Matcher
	callbacks map[GroupEvent][]GroupCallback
}

func NewGroup(matcher Matcher) Group {
	return &group{
		entities:  make(map[EntityID]Entity),
		cache:     make([]Entity, 0),
		matcher:   matcher,
		callbacks: make(map[GroupEvent][]GroupCallback),
	}
}

func (g *group) Entities() []Entity {
	if g.cache == nil {
		cache := make([]Entity, len(g.entities))
		i := 0
		for _, e := range g.entities {
			cache[i] = e
			i++
		}
		g.cache = cache
	}
	return g.cache
}

func (g *group) HandleEntity(e Entity) {
	if g.matcher.Matches(e) {
		g.addEntity(e)
	} else {
		g.removeEntity(e)
	}
}

func (g *group) ContainsEntity(e Entity) bool {
	if _, ok := g.entities[e.ID()]; ok {
		return true
	}
	return false
}

func (g *group) AddCallback(ev GroupEvent, c GroupCallback) {
	cs, ok := g.callbacks[ev]
	if !ok {
		cs = make([]GroupCallback, 0)
	}
	g.callbacks[ev] = append(cs, c)
}

func (g *group) addEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; !ok {
		g.entities[e.ID()] = e
		g.cache = append(g.cache, e)
		g.callback(EntityAdded, e)
	}
}

func (g *group) removeEntity(e Entity) {
	if _, ok := g.entities[e.ID()]; ok {
		delete(g.entities, e.ID())
		if i := findIndex(g.cache, e); i != -1 {
			g.cache = nil
		}
		g.callback(EntityRemoved, e)
	}
}

func (g *group) callback(ev GroupEvent, e Entity) {
	if cs, ok := g.callbacks[ev]; ok {
		for _, c := range cs {
			c(g, e)
		}
	}
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
	new := entities[:len(entities)-1]
	return new
}
