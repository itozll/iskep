package build

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/itozll/iskep/pkg/itemplate"
	"github.com/itozll/iskep/pkg/model"
	"github.com/itozll/iskep/pkg/variable/expand"
)

type Argument struct {
	cfg []*model.Argument
}

func NewArgument(cfg []*model.Argument) (*Argument, error) {
	arg := &Argument{
		cfg: cfg,
	}

	err := arg.validate()
	if err != nil {
		return nil, err
	}

	return arg, nil
}

func (arg *Argument) validate() error {
	m := map[string]bool{}
	for _, cfg := range arg.cfg {
		if cfg.Disabled || cfg.Fixed {
			continue
		}

		if cfg.Name == "" {
			return fmt.Errorf("argument: missing name")
		}

		if cfg.Description == "" {
			return fmt.Errorf("argument: missing description")
		}

		// 可变参数需要保持唯一
		if m[cfg.Name] {
			return fmt.Errorf("argument: duplicate name (%s)", cfg.Name)
		}
		m[cfg.Name] = true
	}

	return nil
}

func (arg *Argument) Questions() []*survey.Question {
	var questions []*survey.Question

	for _, cfg := range arg.cfg {
		if cfg.Disabled || cfg.Fixed {
			continue
		}

		if len(cfg.Range) > 0 {
			questions = append(questions, &survey.Question{
				Name: cfg.Name,
				Prompt: &survey.Select{
					Message: cfg.Description,
					Options: cfg.Range,
					Default: cfg.Value,
				},
			})
		} else {
			prompt := &survey.Question{
				Name: cfg.Name,
				Prompt: &survey.Input{
					Message: cfg.Description,
					Default: cfg.Value,
				},
			}

			if cfg.Required {
				prompt.Validate = survey.Required
			}

			questions = append(questions, prompt)
		}
	}

	return questions
}

func (arg *Argument) CompleteInterface(mapping map[string]any) {
	m := map[string]string{}

	for key, val := range mapping {
		switch val := val.(type) {
		case string:
			m[key] = val
		case core.OptionAnswer:
			m[key] = val.Value
		}
	}

	_ = arg.Complete(m)
}

func (arg *Argument) Complete(mapping map[string]string) (err error) {
	for _, cfg := range arg.cfg {
		if cfg.Disabled || cfg.Fixed {
			continue
		}

		if v, ok := mapping[cfg.Name]; ok {
			cfg.Value = v
			continue
		}

		if cfg.Value == "" {
			return fmt.Errorf("argument: missing value (%s)", cfg.Name)
		}
	}

	return
}

func (arg *Argument) GenerateArguments() map[string]string {
	arguments := map[string]string{}
	fixed := map[string]bool{}
	for _, cfg := range arg.cfg {
		if !cfg.Fixed && fixed[cfg.Name] {
			continue
		}

		value, err := itemplate.Parse(cfg.Value, arguments)
		if err != nil {
			return nil
		}

		if !cfg.Expansion {
			arguments[cfg.Name] = value
		} else {
			for key, value := range expand.Do(cfg.Name, value) {
				arguments[key] = value
			}
		}

		if cfg.Fixed {
			fixed[cfg.Name] = true
		}
	}

	return arguments
}
