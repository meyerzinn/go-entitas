package entitas

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroup(t *testing.T) {

	Convey("Given a new group", t, func() {
		g := NewGroup(AllOf([]ComponentType{}))

		Convey("It gets empty group for matcher when no entities were created", func() {
			So(g.Entities(), ShouldBeEmpty)
		})

		Convey("It should be empty", func() {
			So(g.Count(), ShouldEqual, 0)
		})

		Convey("When entity is added", func() {
			e1 := NewEntity(0, IndexLength)
			g.HandleEntity(e1)

			Convey("It should contain the matching entity", func() {
				So(g.Entities(), ShouldContain, e1)
			})
		})

	})

}
