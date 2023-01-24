package expand

func Snake(str string) string {
	return snake(split(str))
}

func snake(elements []string) string {
	return _separate(elements, "_")
}
