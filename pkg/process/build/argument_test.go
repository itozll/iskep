package build

import (
	"testing"

	"github.com/itozll/iskep/pkg/model"
	"github.com/smartystreets/goconvey/convey"
)

func TestArgument(t *testing.T) {
	cfg := []*model.Argument{
		{
			Name:        "name",
			Description: "name",
			Value:       "hello",
		},
		{
			Name:        "name",
			Description: "name",
			Fixed:       true,
			Value:       "{{ .name }}-world",
		},
		{
			Name:        "struct",
			Description: "struct",
			Value:       "hello-world-v1",
			Expansion:   true,
		},
		{
			Name:        "user",
			Description: "user",
			Required:    true,
		},
		{
			Name:        "fixed",
			Description: "fixed",
			Fixed:       true,
			Value:       "value",
		},
		{
			Name:        "fixed",
			Description: "fixed",
			Value:       "value_1",
		},
		{
			Name:        "fixed",
			Description: "fixed",
			Fixed:       true,
			Value:       "value_3",
		},
	}
	convey.Convey("test argument", t, func() {
		arguments, err := NewArgument(cfg)
		convey.So(err, convey.ShouldBeNil)
		res := arguments.GenerateArguments()
		convey.So(res["name"], convey.ShouldEqual, "hello-world")
		convey.So(res["struct_pascal"], convey.ShouldEqual, "HelloWorldV1")
		convey.So(res["struct_snake"], convey.ShouldEqual, "hello_world_v1")
		convey.So(res["user"], convey.ShouldEqual, "")

		err = arguments.Complete(map[string]string{"user": "itozll"})
		convey.So(res["user"], convey.ShouldEqual, "")
		convey.So(err, convey.ShouldBeNil)

		convey.So(res["fixed"], convey.ShouldEqual, "value_3")
	})
}
