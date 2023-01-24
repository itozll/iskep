package expand

func Separate(str, delimiter string) string {
	return _separate(split(str), delimiter)
}
