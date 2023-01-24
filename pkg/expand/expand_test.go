package expand_test

import (
	"testing"

	"github.com/itozll/iskep/pkg/expand"
	"github.com/smartystreets/goconvey/convey"
)

var test_elements = []struct {
	str    string
	pascal string
	camel  string
	snake  string
	kebad  string
}{
	{"hello_http_v1", "HelloHTTPV1", "helloHTTPV1", "hello_http_v1", "hello-hello-v1"},
	{"hello_httpv1", "HelloHTTPV1", "helloHTTPV1", "hello_http_v1", "hello-hello-v1"},
	{"HTTPhelloV123", "HTTPHelloV123", "httpHelloV123", "http_hello_v123", "http-hello-v123"},
	{"HTTPhellop123", "HTTPHellop123", "httpHellop123", "http_hellop_123", "http-hellop-123"},
	{"HTTP###@hellop123", "HTTPHellop123", "httpHellop123", "http_hellop_123", "http-hellop-123"},
	{"HTTP hellop123", "HTTPHellop123", "httpHellop123", "http_hellop_123", "http-hellop-123"},
}

func TestExpand(t *testing.T) {
	convey.Convey("Test expand", t, func() {
		for _, elem := range test_elements {
			convey.Convey("Test `"+elem.str+"'", func() {
				convey.ShouldEqual(expand.Pascal(elem.str), elem.pascal)
				convey.ShouldEqual(expand.Camel(elem.str), elem.camel)
				convey.ShouldEqual(expand.Snake(elem.str), elem.snake)
				convey.ShouldEqual(expand.Kebad(elem.str), elem.kebad)
			})
		}
	})
}
