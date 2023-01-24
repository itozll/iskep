package expand

import (
	"strings"
	"unicode"
)

type chType int8

const (
	chUpper chType = iota
	chLower
	chDigit
	chUnknown
)

func hitCharType(c rune) chType {
	switch {
	case unicode.IsUpper(c):
		return chUpper
	case unicode.IsLower(c):
		return chLower
	case unicode.IsDigit(c):
		return chDigit
	}

	return chUnknown
}

func split(name string) (v []string) {
	if len(name) == 0 {
		return
	}

	var ch rune
	curStr := ""
	typ := hitCharType(rune(name[1]))

	for _, c := range name {
		if int32(c) > unicode.MaxLatin1 {
			if curStr != "" {
				v = _append(v, curStr)
			}

			curStr = ""
			typ = chUnknown
			ch = 0
			continue
		}

		t := hitCharType(c)
		if t == chUnknown {
			if curStr != "" {
				v = _append(v, curStr)
				curStr = ""
			}

			typ = t
			ch = 0
			continue
		}

		if typ == t {
			curStr += string(c)
			ch = c
			continue
		}

		switch {
		case t == chDigit && (ch == 'v' || ch == 'V'), t == chLower && typ == chUpper:
			if len(curStr) > 1 {
				if FixedElements[strings.ToLower(curStr)] {
					v = _append(v, curStr)
					curStr = ""
				} else {
					v = _append(v, curStr[:len(curStr)-1])
					curStr = string(ch)
				}
			}

		default:
			if curStr != "" {
				v = _append(v, curStr)
				curStr = ""
			}
		}

		curStr += string(c)
		typ = t
		ch = c
	}

	if curStr != "" {
		v = _append(v, curStr)
	}

	return
}

func _append(v []string, str string) []string {
	return append(v, strings.ToLower(str))
}
