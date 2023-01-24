package expand

func Pascal(str string) string {
	return pascal(split(str))
}

func pascal(elements []string) string {
	return _pascal(elements, false)
}
