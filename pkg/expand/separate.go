package expand

// Snake 转换字符串格式，调用者指定分隔符
//
//	HelloWorld  -> hello{delimiter}world
//	Hello_World -> hello{delimiter}world
//	Hello World -> hello{delimiter}world
//	hello-world -> hello{delimiter}world
func Separate(str, delimiter string) string {
	return _separate(split(str), delimiter)
}
