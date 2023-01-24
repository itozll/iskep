package expand

import "strings"

var FixedString = []string{"json", "jsonp", "id", "http", "rpc"}

//	var FixedElements = map[string]bool{
//		"json":  true,
//		"jsonp": true,
//		"id":    true,
//		"http":  true,
//		"rpc":   true,
//	}
var FixedElements = GenFixedElements(FixedString)

func GenFixedElements(fixed []string, strs ...string) map[string]bool {
	v := map[string]bool{}
	for _, fix := range append(fixed, strs...) {
		v[strings.ToLower(fix)] = true
	}
	return v
}
