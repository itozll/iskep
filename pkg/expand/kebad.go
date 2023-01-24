package expand

func Kebad(str string) string {
	return kebad(split(str))
}

func kebad(elements []string) string {
	return _separate(elements, "_")
}
