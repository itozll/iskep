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
