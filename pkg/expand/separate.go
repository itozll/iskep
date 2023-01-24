package expand

import "strings"

func Snake(str string) string {
	return snake(split(str))
}

func snake(elements []string) string {
	return separate(elements, "_")
}

func Kebad(str string) string {
	return kebad(split(str))
}

func kebad(elements []string) string {
	return separate(elements, "_")
}

func Separate(str, delimiter string) string {
	return separate(split(str), delimiter)
}

func separate(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}
