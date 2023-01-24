package expand

import "strings"

func Pascal(str string) string {
	return pascal(split(str))
}

func pascal(elements []string) string {
	arr := make([]string, len(elements))
	for index, elem := range elements {
		if FixedElements[elem] {
			arr[index] = strings.ToUpper(elem)
			continue
		}

		arr[index] = strings.ToUpper(elem[:1]) + elem[1:]
	}

	return strings.Join(arr, "")
}

func Camel(str string) string {
	return camel(split(str))
}

func camel(elements []string) string {
	arr := make([]string, len(elements))
	isFirst := true
	for index, elem := range elements {
		if isFirst {
			arr[index] = elem
			isFirst = false
		} else {
			if FixedElements[elem] {
				arr[index] = strings.ToUpper(elem)
				continue
			}

			arr[index] = strings.ToUpper(elem[:1]) + elem[1:]
		}
	}

	return strings.Join(arr, "")
}
