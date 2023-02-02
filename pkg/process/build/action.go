package build

import (
	"errors"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/itozll/iskep/pkg/model"
	"github.com/itozll/iskep/pkg/process"
	"github.com/itozll/iskep/pkg/runtime/rtstatus"
	"github.com/itozll/iskep/pkg/tmpl"
)

// Action 命令
type Action struct {
	cfg *model.Action

	before func() error
	after  func() error
	p      tmpl.Provider

	// // 当前脚本参数列表
	// Arguments Argument

	// // 子命令列表
	// Actions []*Action
}

func NewAction(cfg *model.Action, cmd *Command) *Action {
	action := &Action{
		cfg:    cfg,
		before: process.Commands(cfg.Before),
		after:  process.Commands(cfg.After),
	}

	if cfg.Provider == "" {
		action.p = cmd.p
	} else {
		action.p = tmpl.GetTemplateProvider(cfg.Provider)
	}

	return action
}

func (action *Action) Exec(path string, binder map[string]string) (err error) {
	if err = action.before(); err != nil {
		return
	}

	if path != "" {
		path = strings.TrimRight(path, "/") + "/"
	}

	if len(action.cfg.To) > 0 {
		path += strings.TrimRight(action.cfg.To, "/") + "/"
	}

	if len(action.cfg.Template) > 0 {
		if err := action.parseAndCopy(path, action.cfg.Template, binder, true); err != nil {
			return err
		}
	}

	if len(action.cfg.Copy) > 0 {
		if err := action.parseAndCopy(path, action.cfg.Copy, binder, false); err != nil {
			return err
		}
	}

	return action.after()
}

func (action *Action) parseAndCopy(path string, list []string, binder map[string]string, isTmpl bool) error {
	if len(path) > 0 {
		if err := os.MkdirAll(path, os.ModePerm); err != nil && os.IsNotExist(err) {
			rtstatus.Error("%s (%s)", err, path)
			return err
		}
	}

	for _, name := range list {
		if len(name) == 0 {
			return errors.New("source name must not be empty")
		}

		dstName, srcName := splitName(name)
		dstPath := path + dstName

		if err := func() error {
			dstFd, err := os.Create(dstPath)
			if err != nil {
				return err
			}
			defer dstFd.Close()

			data := action.p.ReadFile(srcName)
			if !isTmpl {
				_, err = io.Copy(dstFd, strings.NewReader(string(data)))
				if err != nil {
					return err
				}
			} else {
				tpl, err := template.New(dstPath).Parse(string(data))
				if err != nil {
					return err
				}

				err = tpl.Execute(dstFd, binder)
				if err != nil {
					return err
				}
			}

			return nil
		}(); err != nil {
			return err
		}

		rtstatus.Info("Create", dstPath)
	}

	return nil
}

func splitName(name string) (dst, src string) {
	names := strings.Split(name, ":")
	src = names[0]
	dst = src

	l := len(names)

	// name
	if l == 1 {
		return
	}

	if l > 2 {
		// name::new-name
		if len(names[2]) > 0 {
			dst = names[2]
			return
		}
	}

	// name:<suffix> -> model:go change to model.go
	if len(names[1]) > 0 {
		dst += names[1]
	}

	return
}
