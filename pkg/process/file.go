package process

import (
	"os"

	"github.com/itozll/iskep/internal/etcd"
	"github.com/itozll/iskep/pkg/runtime/rtstatus"
)

func Chdir(path string) error {
	return os.Chdir(path)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func PathNotExists(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

func ReadFile(name string) []byte {
	_, err := os.Stat(name)
	rtstatus.ExitIfError(err)

	data, err := os.ReadFile(name)
	rtstatus.ExitIfError(err)

	return data
}

var Expand = etcd.Expand
