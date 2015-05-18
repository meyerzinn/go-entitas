package entitas

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// Component index

const (
	IndexComponent1 ComponentType = iota
	IndexComponent2
	IndexComponent3
	IndexLength
)

// Component 1

type component1 struct {
	value int
}

func NewComponent1(value int) Component {
	return &component1{value: value}
}

func (c1 *component1) Type() ComponentType {
	return IndexComponent1
}

func (c1 *component1) String() string {
	return "Component1"
}

// Component 2

type component2 struct {
	value float32
}

func NewComponent2(value float32) Component {
	return &component2{value: value}
}

func (c2 *component2) Type() ComponentType {
	return IndexComponent2
}

func (c2 *component2) String() string {
	return "Component2"
}

// Component 3

type component3 struct{}

func NewComponent3() Component {
	return &component3{}
}

func (c *component3) Type() ComponentType {
	return IndexComponent3
}

// Tests

func TestComponentSorting(t *testing.T) {
	Convey("Given components and a component list", t, func() {
		c1 := NewComponent1(1)
		c2 := NewComponent2(0.0)
		components := []Component{c2, c1}

		Convey("It should be sortable by type", func() {
			sort.Sort(ComponentsByType(components))
			So(components, ShouldResemble, []Component{c1, c2})
		})

	})
}
