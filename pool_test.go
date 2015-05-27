package entitas

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPool(t *testing.T) {

	Convey("Given a new pool", t, func() {
		p := NewPool(IndexLength, 0)

		Convey("It has no entities when no entities were created", func() {
			So(len(p.Entities()), ShouldEqual, 0)
		})

		Convey("It creates entity", func() {
			So(p.CreateEntity(), ShouldHaveSameTypeAs, NewEntity(-1))
		})

		Convey("It has a total entity count of 0", func() {
			So(p.Count(), ShouldEqual, 0)
		})

		Convey("It doesn't have entities that were not created with CreateEntity()", func() {
			So(p.HasEntity(NewEntity(-1)), ShouldBeFalse)
		})

		Convey("It gets empty group for matcher when no entities were created", func() {
			g := p.Group(AllOf(IndexComponent1))
			So(g.Entities(), ShouldBeEmpty)
		})

		Convey("It should panic when trying to destroy an entity which doesn't exist", func() {
			e := NewEntity(-1)
			So(func() { p.DestroyEntity(e) }, ShouldPanicWith, "unknown entity")
		})

		Convey("When an entity is created", func() {
			c1 := NewComponent1(1)
			e1 := p.CreateEntity(c1)

			Convey("It has entities that were created", func() {
				So(p.HasEntity(e1), ShouldBeTrue)
			})

			Convey("It has a total entity count of 1", func() {
				So(p.Count(), ShouldEqual, 1)
			})

			Convey("It increments ID", func() {
				So(e1.ID(), ShouldEqual, 0)
			})

			Convey("It destroys entity and removes it", func() {
				e1.AddComponent(c1)
				p.DestroyEntity(e1)
				So(p.HasEntity(e1), ShouldBeFalse)
				So(p.Entities(), ShouldNotContain, e1)
			})

			Convey("When another entity is created", func() {
				e2 := p.CreateEntity()

				Convey("It has a total entity count of 2", func() {
					So(p.Count(), ShouldEqual, 2)
				})

				Convey("It increments ID", func() {
					So(e2.ID(), ShouldEqual, 1)
				})

				Convey("It should have the entity", func() {
					So(p.HasEntity(e2), ShouldBeTrue)
				})

				Convey("It returns all created entities", func() {
					entities := p.Entities()
					So(entities, ShouldContain, e1)
					So(entities, ShouldContain, e2)
				})

				Convey("It can be printed", func() {
					So(fmt.Sprintf("%v", p), ShouldContainSubstring, "Pool(")
					So(fmt.Sprintf("%v", p), ShouldContainSubstring, "Entity_0([Component1])")
					So(fmt.Sprintf("%v", p), ShouldContainSubstring, "Entity_1([])")
				})

				Convey("It should remove that entity when destroyed", func() {
					p.DestroyEntity(e2)
					So(p.HasEntity(e2), ShouldBeFalse)
				})

				Convey("It destroys all entities and removes their components", func() {
					c1 := NewComponent1(9)
					c2 := NewComponent2(3.0)
					e1.AddComponent(c1)
					e2.AddComponent(c2)
					p.DestroyAllEntities()
					So(p.Entities(), ShouldBeEmpty)
					So(e1.Components(), ShouldBeEmpty)
					So(e2.Components(), ShouldBeEmpty)
				})

			})

			Convey("When a group is created", func() {
				g := p.Group(AllOf())

				Convey("The entity should be in the group", func() {
					So(g.Entities(), ShouldContain, e1)
				})
			})

		})

	})

}

func TestPoolEntityID(t *testing.T) {
	Convey("Given a new pool with a different ID", t, func() {
		op := NewPool(IndexLength, 7)

		Convey("It creates new entities with that ID", func() {
			e := op.CreateEntity()
			So(e.ID(), ShouldEqual, 7)
		})
	})
}

func BenchmarkCreateGroup(b *testing.B) {
	p := NewPool(IndexLength, 0)

	for i := 0; i < 2000; i++ {
		p.CreateEntity(
			NewComponent1(i),
			NewComponent2(float32(i)),
			NewComponent3(),
		)
	}

	for n := 0; n < b.N; n++ {
		g := p.Group(AllOf(IndexComponent1))
		g.Entities()
	}
}
