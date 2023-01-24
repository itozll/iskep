package expand

// Kebad 转换字符串格式，短横线分隔
//
//	HelloWorld  -> hello-world
//	Hello_World -> hello-world
//	Hello World -> hello-world
//	hello-world -> hello-world
func Kebad(str string) string {
	return kebad(split(str))
}

func kebad(elements []string) string {
	return _separate(elements, "_")
}
