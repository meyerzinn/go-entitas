package entitas

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEntity(t *testing.T) {

	Convey("When given a new entity", t, func() {
		e := NewEntity(0, IndexLength)
		c1 := NewComponent1(1)
		c2 := NewComponent2(2.0)
		types := []ComponentType{c1.Type(), c2.Type()}

		Convey("It has component of type when component of that type was added", func() {
			e.AddComponent(c1)
			So(e.HasComponent(c1.Type()), ShouldBeTrue)
		})

		Convey("It doesn't have component of type when no component of that type was added", func() {
			So(e.HasComponent(c1.Type()), ShouldBeFalse)
		})

		Convey("It doesn't have components of types when no components of these types were added", func() {
			So(e.HasComponent([]ComponentType{c1.Type()}...), ShouldBeFalse)
		})

		Convey("It doesn't have components of types when not all components of these types were added", func() {
			e.AddComponent(c1)
			So(e.HasComponent(types...), ShouldBeFalse)
		})

		Convey("It has components of types when all components of these types were added", func() {
			e.AddComponent(c1, c2)
			So(e.HasComponent(types...), ShouldBeTrue)
		})

		Convey("It doesn't have any components of types when no components of these types were added", func() {
			So(e.HasAnyComponent(types...), ShouldBeFalse)
		})

		Convey("It has any components of types when any component of these types were added", func() {
			e.AddComponent(c1)
			So(e.HasAnyComponent(types...), ShouldBeTrue)
		})

		Convey("It removes a component of type", func() {
			e.AddComponent(c1)
			e.RemoveComponent(c1.Type())
			So(e.HasComponent(c1.Type()), ShouldBeFalse)
			So(e.Components(), ShouldBeEmpty)
		})

		Convey("It gets a component of type", func() {
			e.AddComponent(c2)
			So(e.Components(), ShouldResemble, []Component{c2})
			c, err := e.Component(c2.Type())
			So(c, ShouldEqual, c2)
			So(err, ShouldBeNil)
		})

		Convey("It doesn't get a component of type that wasn't added", func() {
			c, err := e.Component(c1.Type())
			So(c, ShouldBeNil)
			So(err.Error(), ShouldEqual, "component not found")
		})

		Convey("It replaces an existing component", func() {
			e.AddComponent(c1)
			c11 := NewComponent1(2)
			e.ReplaceComponent(c11)
			actual, err := e.Component(c1.Type())
			So(err, ShouldBeNil)
			So(actual, ShouldNotEqual, c1)
			So(actual, ShouldEqual, c11)
		})

		Convey("It adds a component when replacing a non existing component", func() {
			e.ReplaceComponent(c1)
			c, err := e.Component(c1.Type())
			So(err, ShouldBeNil)
			So(c, ShouldEqual, c1)
		})

		Convey("It returns an empty array of components when no components were added", func() {
			So(e.Components(), ShouldBeEmpty)
		})

		Convey("It returns an empty array of component indices when no components were added", func() {
			So(e.ComponentIndices(), ShouldBeEmpty)
		})

		Convey("It returns all components", func() {
			e.AddComponent(c1, c2)
			So(e.Components(), ShouldContain, c1)
			So(e.Components(), ShouldContain, c2)
		})

		Convey("It returns all component indices", func() {
			e.AddComponent(c1, c2)
			So(e.ComponentIndices(), ShouldContain, c1.Type())
			So(e.ComponentIndices(), ShouldContain, c2.Type())
		})

		Convey("It removes all components", func() {
			e.AddComponent(c1, c2)
			e.RemoveAllComponents()
			actual, err := e.Component(c1.Type())
			So(actual, ShouldBeNil)
			So(err.Error(), ShouldEqual, "component not found")
			actual, err = e.Component(c2.Type())
			So(actual, ShouldBeNil)
			So(err.Error(), ShouldEqual, "component not found")
		})

		Convey("It can be printed", func() {
			e.AddComponent(c1, c2)
			So(fmt.Sprintf("%v", e), ShouldEqual, "Entity_0([Component1 Component2])")
		})

	})

}
