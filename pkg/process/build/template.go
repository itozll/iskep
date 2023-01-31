package build

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

var funcMap = map[string]any{
	"date":          time.Now().Format,
	"has_prefix":    strings.HasPrefix,
	"has_no_prefix": func(s, prefix string) bool { return !strings.HasPrefix(s, prefix) },
	"has_suffix":    strings.HasSuffix,
	"has_no_suffix": func(s, suffix string) bool { return !strings.HasSuffix(s, suffix) },
	"trim":          strings.TrimSpace,
	"trim_prefix":   strings.TrimPrefix,
	"trim_suffix":   strings.TrimSuffix,
	"contain":       strings.Contains,
}

func DoTemplate(tpl string, args map[string]string) (string, error) {
	if !strings.Contains(tpl, "{{") {
		return tpl, nil
	}

	t := template.New("build")
	t.Funcs(funcMap)
	_, err := t.Parse(tpl)
	if err != nil {
		return "", nil
	}

	var result bytes.Buffer
	err = t.Execute(&result, args)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
