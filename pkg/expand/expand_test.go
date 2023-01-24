package expand_test

import (
	"testing"

	"github.com/itozll/iskep/pkg/expand"
)

var test_elements = []string{
	"hello_http_v1",
	"hello_httpv1",
	"HELLOHttpV1",
	"HTTPhelloV123",
	"HTTPhellop123",
	"Ab!@#!@$!@$#&word",
	"",
}

func TestExpand(t *testing.T) {
	for _, elem := range test_elements {
		t.Log(elem, expand.Pascal(elem), expand.Snake(elem), expand.Camel(elem))
	}
}
