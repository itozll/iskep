package itemplate

import (
	"bytes"
	"strings"
	"text/template"
	"time"
)

var DelimLeft = "{{"
var DelimRight = "}}"

var FuncMap = map[string]any{
	"date":          time.Now().Format,
	"has_prefix":    strings.HasPrefix,
	"has_no_prefix": func(s, prefix string) bool { return !strings.HasPrefix(s, prefix) },
	"has_suffix":    strings.HasSuffix,
	"has_no_suffix": func(s, suffix string) bool { return !strings.HasSuffix(s, suffix) },
	"trim":          strings.TrimSpace,
	"trim_prefix":   strings.TrimPrefix,
	"trim_suffix":   strings.TrimSuffix,
	"contain":       strings.Contains,
	"not_contain":   func(s, substr string) bool { return !strings.Contains(s, substr) },
	"to_upper":      strings.ToUpper,
	"to_lower":      strings.ToLower,
	"replace":       strings.ReplaceAll,
	"split":         strings.Split,
}

func Parse(tpl string, args map[string]string) (string, error) {
	if !strings.Contains(tpl, "{{") || !strings.Contains(tpl, "}}") {
		return tpl, nil
	}

	t, err := template.New("itemplate").
		Delims(DelimLeft, DelimRight).
		Funcs(FuncMap).
		Parse(tpl)
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
