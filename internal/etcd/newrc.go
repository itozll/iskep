package etcd

import "github.com/itozll/iskep/pkg/model"

var NewCommandConfig = &model.Command{
	Actions: []*model.Action{
		{
			Template: []string{
				"README:.md",
				"Makefile",
				"main_go::main.go",
				"gomod::go.mod",
			},
			Copy: []string{
				"gitignore::.gitignore",
				"generator:.sh",
			},
		},
		{
			To: "app/cmd",
			Template: []string{
				"cmdroot::root.go",
				"cmdserver::server.go",
			},
		},
		{
			To: "app/internal/runtime/rtinfo",
			Template: []string{
				"context_go::context.go",
			},
			Copy: []string{
				"rtinfo_go::rtinfo.go",
			},
		},
	},
}
