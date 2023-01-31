package build

import (
	"fmt"

	"github.com/itozll/iskep/pkg/model"
	"github.com/itozll/iskep/pkg/variable/expand"
)

type Argument struct {
	cfg   []*model.Argument
	v     map[string]string
	fixed map[string]bool

	// 声明了 required 但不包含值的参数集合
	lost map[string]*model.Argument
}

func NewArgument(cfg []*model.Argument) (*Argument, error) {
	arg := &Argument{
		cfg:   cfg,
		v:     make(map[string]string),
		fixed: make(map[string]bool),
		lost:  make(map[string]*model.Argument),
	}

	for _, v := range cfg {
		if v.Disabled {
			continue
		}

		if v.Name == "" {
			return nil, fmt.Errorf("argument: missing name")
		}

		// 当参数设置为 fixed 时，禁止替换该值
		// 这有一个前提，字面值为空时允许替换
		if arg.fixed[v.Name] {
			continue
		}

		if v.Value == "" {
			if v.Required {
				arg.lost[v.Name] = v
			}

			continue
		}

		// 保存该参数的 fixed 状态
		if v.Fixed {
			arg.fixed[v.Name] = true
		}

		value, err := DoTemplate(v.Value, arg.v)
		if err != nil {
			return nil, err
		}

		if !v.Expansion {
			arg.v[v.Name] = value
			continue
		}

		for key, value := range expand.Do(v.Name, value) {
			arg.v[key] = value
		}
	}

	return arg, nil
}

func (arg *Argument) Arguments() map[string]string { return arg.v }

func (arg *Argument) Load(cfg map[string]string) (args []*model.Argument) {
loop:
	for key, value := range cfg {
		if value == "" {
			continue
		}

		delete(arg.lost, key)

		for _, orig := range arg.cfg {
			if key == orig.Name && orig.Expansion {
				for k, v := range expand.Do(key, value) {
					arg.v[k] = v
				}

				continue loop
			}
		}

		arg.v[key] = value
	}

	for _, v := range arg.lost {
		args = append(args, v)
	}

	return
}
