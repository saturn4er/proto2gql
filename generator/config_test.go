package generator

import (
	"testing"
)

func TestProtoConfig_Copy(t *testing.T) {
	Convey("Should copy ProtoConfig", t, func() {
		So(func() { ProtoConfig{}.Copy() }, ShouldNotPanic)
	})
}
