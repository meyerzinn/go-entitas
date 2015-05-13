package entitas

type Matcher interface {
	Matches(entity Entity) bool
}

type AllOf []ComponentType

func (allof AllOf) Matches(entity Entity) bool {
	return entity.HasComponent(allof...)
}
