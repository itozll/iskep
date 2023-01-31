package tmpl

import "github.com/itozll/iskep/internal/template"

func init() {
	AddTemplateProvider("base", template.Builtin())
}
