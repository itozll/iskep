package etcd

import (
	"embed"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed *.yaml
var fs embed.FS

func Unmarshal(name string, mapping map[string]string, out any) error {
	if !strings.HasSuffix(name, "yaml") {
		name += ".yaml"
	}

	data, err := fs.ReadFile(name)
	if err != nil {
		return err
	}

	return yaml.Unmarshal([]byte(Expand(data, mapping)), out)
}

func Expand(data []byte, mapping map[string]string) string {
	str := string(data)

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
