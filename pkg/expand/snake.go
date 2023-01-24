package expand

// Snake 转换字符串格式，下划线分隔
//
//	HelloWorld  -> hello_world
//	Hello_World -> hello_world
//	Hello World -> hello_world
//	hello-world -> hello_world
func Snake(str string) string {
	return snake(split(str))
}

func snake(elements []string) string {
	return _separate(elements, "_")
}
