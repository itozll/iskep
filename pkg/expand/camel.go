package expand

// Camel 转换字符串格式，小驼峰
//
//	HelloWorld  -> helloWorld
//	Hello_World -> helloWorld
//	Hello World -> helloWorld
//	hello-world -> helloWorld
func Camel(str string) string {
	return camel(split(str))
}

func camel(elements []string) string {
	return _pascal(elements, true)
}
