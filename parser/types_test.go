package parser

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeIs(t *testing.T) {
	Convey("Test Type", t, func() {
		Convey("Type is Message", func() {
			msgTyp := &Type{Message: &Message{}}
			So(msgTyp.IsMessage(), ShouldBeTrue)
			So(msgTyp.IsMap(), ShouldBeFalse)
			So(msgTyp.IsEnum(), ShouldBeFalse)
			So(msgTyp.IsScalar(), ShouldBeFalse)
		})

		Convey("Type is Enum", func() {
			enumTyp := &Type{Enum: &Enum{}}
			So(enumTyp.IsEnum(), ShouldBeTrue)
			So(enumTyp.IsMap(), ShouldBeFalse)
			So(enumTyp.IsMessage(), ShouldBeFalse)
			So(enumTyp.IsScalar(), ShouldBeFalse)
		})

		Convey("Type is Map", func() {
			enumTyp := &Type{Map: &Map{}}
			So(enumTyp.IsMap(), ShouldBeTrue)
			So(enumTyp.IsEnum(), ShouldBeFalse)
			So(enumTyp.IsMessage(), ShouldBeFalse)
			So(enumTyp.IsScalar(), ShouldBeFalse)
		})

		Convey("Type is Scalar", func() {
			enumTyp := &Type{Scalar: "int"}
			So(enumTyp.IsScalar(), ShouldBeTrue)
			So(enumTyp.IsMap(), ShouldBeFalse)
			So(enumTyp.IsEnum(), ShouldBeFalse)
			So(enumTyp.IsMessage(), ShouldBeFalse)
		})
	})
}
