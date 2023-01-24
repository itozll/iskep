package expand

// Pascal 转换字符串格式，驼峰
//
//	HelloWorld  -> HelloWorld
//	Hello_World -> HelloWorld
//	Hello World -> HelloWorld
//	hello-world -> HelloWorld
func Pascal(str string) string {
	return pascal(split(str))
}

func pascal(elements []string) string {
	return _pascal(elements, false)
}
