package build

import (
	"fmt"

	"github.com/itozll/iskep/pkg/model"
	"github.com/itozll/iskep/pkg/process"
	"github.com/itozll/iskep/pkg/runtime/rtinfo"
	"github.com/itozll/iskep/pkg/tmpl"
)

type Command struct {
	cfg *model.Command

	before func() error
	after  func() error
	p      tmpl.Provider

	binder    map[string]string
	arguments *Argument

	actions []*Action
}

func NewCommand(cfg *model.Command) (*Command, error) {
	if len(cfg.Actions) == 0 {
		return nil, fmt.Errorf("no action")
	}

	p := tmpl.GetTemplateProvider(cfg.Provider, "base")
	if p == nil {
		return nil, fmt.Errorf("no template provider")
	}

	arguments, err := NewArgument(cfg.Arguments)
	if err != nil {
		return nil, err
	}

	c := &Command{
		cfg:       cfg,
		p:         p,
		binder:    make(map[string]string),
		before:    process.Commands(cfg.Before),
		after:     process.Commands(cfg.After),
		arguments: arguments,
		actions:   make([]*Action, len(cfg.Actions)),
	}

	for index, action := range cfg.Actions {
		c.actions[index] = NewAction(action, c)
	}

	return c, nil
}

func (cmd *Command) Attach(key, value string) { cmd.binder[key] = value }
func (cmd *Command) AttachMap(binder map[string]string) {
	for key, value := range binder {
		cmd.Attach(key, value)
	}
}

func (cmd *Command) Exec(args map[string]string) (err error) {
	lost := cmd.arguments.Load(args)
	if lost != nil {
		return fmt.Errorf("lost argument")
	}

	if err = cmd.before(); err != nil {
		return
	}

	for key, value := range cmd.arguments.Arguments() {
		cmd.binder[key] = value
	}

	for _, action := range cmd.actions {
		if err = action.Exec(rtinfo.TargetPath, cmd.binder); err != nil {
			return err
		}
	}

	return cmd.after()
}
