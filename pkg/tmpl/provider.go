package tmpl

import (
	"github.com/itozll/iskep/internal/template"
	"github.com/itozll/iskep/pkg/runtime/rtstatus"
)

type Provider = template.Provider

var (
	WithFS    = template.WithFS
	providers = map[string]Provider{}
)

func GetTemplateProvider(name string, base ...string) Provider {
	if name == "" && len(base) > 0 {
		name = base[0]
	}

	v, ok := providers[name]
	if !ok {
		rtstatus.Fatal("no such template named '" + name + "'")
	}

	return v
}

func AddTemplateProvider(name string, p Provider) {
	providers[name] = p
}
