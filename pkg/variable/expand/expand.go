package expand

import (
	"strings"
	"unicode"
)

func Do(name, value string) map[string]string {
	elements := split(value)

	return map[string]string{
		name + "_pascal": pascal(elements),
		name + "_camel":  camel(elements),
		name + "_snake":  snake(elements),
		name + "_kebad":  kebad(elements),
	}
}

type chType int8

const (
	chInvalid chType = iota
	chUpper
	chLower
	chDigit
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

	return chInvalid
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
			typ = chInvalid
			ch = 0
			continue
		}

		t := hitCharType(c)
		if t == chInvalid {
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

func _separate(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}

func _pascal(elements []string, ignoreTheFirst bool) string {
	arr := make([]string, len(elements))
	for index, elem := range elements {
		if ignoreTheFirst {
			arr[index] = elem
			ignoreTheFirst = false
		} else {
			if FixedElements[elem] {
				arr[index] = strings.ToUpper(elem)
				continue
			}

			arr[index] = strings.ToUpper(elem[:1]) + elem[1:]
		}
	}

	return strings.Join(arr, "")
}
