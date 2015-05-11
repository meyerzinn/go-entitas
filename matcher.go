package entitas

type Matcher interface {
	Matches(entity Entity) bool
}

type AllOf []ComponentType

func (allof AllOf) Matches(entity Entity) bool {
	if len(entity.Components()) == 0 {
		return false
	}

	for _, c := range entity.Components() {
		for _, t := range allof {
			if c.Type() != t {
				return false
			}
		}
	}
	return true
}
