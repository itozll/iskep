package template

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"strings"

	"github.com/itozll/iskep/pkg/runtime/rtstatus"
)

type Provider interface {
	ReadFile(string) []byte
	ReadDir(string) []fs.DirEntry
}

type provide struct {
	f    embed.FS
	path string
}

func newProvide(path string, f embed.FS) Provider {
	return &provide{path: strings.TrimSuffix(path, "/") + "/", f: f}
}

func (p *provide) ReadFile(name string) []byte {
	data, err := p.get(p.path + name)
	rtstatus.ExitIfError(err)
	return data
}

func (p *provide) ReadDir(name string) []fs.DirEntry {
	de, err := p.f.ReadDir(p.path + name)
	rtstatus.ExitIfError(err)
	return de
}

func (p *provide) get(name string) ([]byte, error) {
	data, err := p.f.ReadFile(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// try <name>.tpl
			if !strings.HasSuffix(name, ".tpl") {
				return p.get(name + ".tpl")
			}
		}

		rtstatus.Fatal("\n  ** can not get template[%s]: %s.", name, err)
	}

	return data, nil
}
