package expand

func Camel(str string) string {
	return camel(split(str))
}

func camel(elements []string) string {
	return _pascal(elements, true)
}
