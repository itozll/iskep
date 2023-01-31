package process

import (
	"os"

	"github.com/itozll/iskep/pkg/runtime/rtstatus"
)

func Chdir(path string) {
	rtstatus.ExitIfError(os.Chdir(path))
}

func ReadFile(name string) []byte {
	_, err := os.Stat(name)
	rtstatus.ExitIfError(err)

	data, err := os.ReadFile(name)
	rtstatus.ExitIfError(err)

	return data
}

func Expand(body []byte, mapping map[string]string) string {
	str := string(body)

	if len(mapping) != 0 {
		str = os.Expand(str, func(s string) string {
			if v, ok := mapping[s]; ok {
				return v
			}

			return "$" + s
		})
	}

	return os.ExpandEnv(str)
}
