package entitas

const (
	MATCHER_HASH_ID_PRIME  = 647
	MATCHER_HASH_LEN_PRIME = 997
)

type Matcher interface {
	Matches(entity Entity) bool
	Hash() MatcherHash
	ComponentTypes() []ComponentType
}

type MatcherHash uint

type BaseMatcher struct {
	ids  []ComponentType
	hash MatcherHash
}

// --- BaseMatcher ------------------------------------------------------------

func NewBaseMatcher(ids ...ComponentType) BaseMatcher {
	b := BaseMatcher{ids: ids}
	b.hash = b.Hash()
	return b
}

func (b *BaseMatcher) Hash() MatcherHash {
	return b.hash
}

func (b *BaseMatcher) ComponentTypes() []ComponentType {
	return b.ids
}

// --- AllOf ------------------------------------------------------------------

type AllMatcher struct {
	BaseMatcher
}

func AllOf(ids ...ComponentType) Matcher {
	b := NewBaseMatcher(ids...)
	b.hash = Hash(ids...)
	return &AllMatcher{b}
}

func (a *AllMatcher) Matches(e Entity) bool {
	return e.HasComponent(a.ids...)
}

// --- AnyOf ------------------------------------------------------------------

type AnyMatcher struct {
	BaseMatcher
}

func AnyOf(ids ...ComponentType) Matcher {
	b := NewBaseMatcher(ids...)
	b.hash = Hash(ids...)
	return &AnyMatcher{b}
}

func (a *AnyMatcher) Matches(e Entity) bool {
	return e.HasAnyComponent(a.ids...)
}

// --- Utilities --------------------------------------------------------------

func Hash(ids ...ComponentType) MatcherHash {
	var hash uint
	for id := range ids {
		hash ^= uint(id) * MATCHER_HASH_ID_PRIME
	}
	hash ^= uint(len(ids)) * MATCHER_HASH_LEN_PRIME
	return MatcherHash(hash)
}
