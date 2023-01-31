package template

import "embed"

//go:embed templates/*
var f embed.FS

func Builtin() Provider {
	return newProvide("templates", f)
}

func WithFS(path string, _f embed.FS) Provider {
	return newProvide(path, _f)
}
