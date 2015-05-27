package entitas

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroup(t *testing.T) {

	Convey("Given a new group matching all of one component", t, func() {
		g := NewGroup(AllOf(IndexComponent1))
		e1 := NewEntity(0)
		e1.AddComponent(NewComponent1(5))

		Convey("It gets empty group for matcher when no entities were created", func() {
			So(g.Entities(), ShouldBeEmpty)
		})

		Convey("It should not contain an entity", func() {
			So(g.ContainsEntity(e1), ShouldBeFalse)
		})

		Convey("When entity is added", func() {
			g.HandleEntity(e1)

			Convey("The entity should be in the group's entities", func() {
				So(g.Entities(), ShouldContain, e1)
			})

			Convey("It is not empty", func() {
				So(len(g.Entities()), ShouldEqual, 1)
			})

			Convey("It should contain the matching entity", func() {
				So(g.ContainsEntity(e1), ShouldBeTrue)
			})

			Convey("It should still exist when added twice", func() {
				g.HandleEntity(e1)
				So(g.ContainsEntity(e1), ShouldBeTrue)
			})
		})

		Convey("When non-matching entity is added", func() {
			e1.RemoveComponent(IndexComponent1)
			g.HandleEntity(e1)

			Convey("The entity should not be in the group's entities", func() {
				So(g.Entities(), ShouldNotContain, e1)
			})
		})

		Convey("When an matching entity component is removed", func() {
			g.HandleEntity(e1)
			e1.RemoveComponent(IndexComponent1)
			g.HandleEntity(e1)

			Convey("The entity should not be in the group's entities", func() {
				So(g.Entities(), ShouldNotContain, e1)
			})
		})

	})

	Convey("Given a new group matching all of two component", t, func() {
		g := NewGroup(AllOf(IndexComponent1, IndexComponent2))
		e1 := NewEntity(0)
		e1.AddComponent(NewComponent1(5))
		e1.AddComponent(NewComponent3())
		e1.AddComponent(NewComponent2(5))

		Convey("When entity is added", func() {
			g.HandleEntity(e1)

			Convey("The entity should be in the group's entities", func() {
				So(g.Entities(), ShouldContain, e1)
			})
		})
	})
}

func BenchmarkGroupHandleEntity(b *testing.B) {
	c1 := NewComponent1(1)
	c2 := NewComponent2(1.0)
	g := NewGroup(AllOf(IndexComponent1, IndexComponent2))

	e := &entity{
		id:         0,
		components: make(map[ComponentType]Component),
		callbacks:  make(map[ComponentEvent][]ComponentCallback),
	}
	Entity(e).AddComponent(c1, c2)
	for n := 0; n < b.N; n++ {
		e.id = EntityID(n % 100) // Add and remove 100 different entities
		g.HandleEntity(e)
	}
}
