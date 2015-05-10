package entitas

type ComponentType int16
type Component interface {
	Type() ComponentType
}

type ComponentsByType []Component

func (t ComponentsByType) Len() int {
	return len(t)
}
func (t ComponentsByType) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ComponentsByType) Less(i, j int) bool {
	return t[i].Type() < t[j].Type()
}
