package build

import (
	"testing"

	"github.com/itozll/iskep/pkg/model"
	"github.com/smartystreets/goconvey/convey"
)

func TestArgument(t *testing.T) {
	cfg := []*model.Argument{
		{
			Name:  "name",
			Value: "hello",
		},
		{
			Name:  "name",
			Value: "{{ .name }}-world",
		},
		{
			Name:      "struct",
			Value:     "hello-world-v1",
			Expansion: true,
		},
		{
			Name:     "user",
			Required: true,
		},
		{
			Name:  "fixed",
			Fixed: true,
			Value: "value",
		},
		{
			Name:  "fixed",
			Value: "value_1",
		},
		{
			Name:  "fixed2",
			Fixed: true,
		},
		{
			Name:  "fixed2",
			Value: "value_1",
		},
	}
	convey.Convey("test argument", t, func() {
		arguments, err := NewArgument(cfg)
		convey.So(err, convey.ShouldBeNil)
		res := arguments.Arguments()
		convey.So(res["name"], convey.ShouldEqual, "hello-world")
		convey.So(res["struct_pascal"], convey.ShouldEqual, "HelloWorldV1")
		convey.So(res["struct_snake"], convey.ShouldEqual, "hello_world_v1")
		convey.So(res["user"], convey.ShouldEqual, "")
		convey.So(len(arguments.lost), convey.ShouldEqual, 1)

		args := arguments.Load(map[string]string{"user": "itozll"})
		convey.So(res["user"], convey.ShouldEqual, "itozll")
		convey.So(args, convey.ShouldBeNil)
		convey.So(len(arguments.lost), convey.ShouldEqual, 0)

		convey.So(res["fixed"], convey.ShouldEqual, "value")
		convey.So(res["fixed2"], convey.ShouldEqual, "value_1")
	})
}
