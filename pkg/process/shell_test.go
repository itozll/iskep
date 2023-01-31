package process

import (
	"os"
	"testing"
)

func TestCommand(t *testing.T) {
	cmdstr := []string{
		`go mod tidy`,
		`echo "hello    world"`,
	}

	for _, str := range cmdstr {
		CommandOutput(str, os.Stdout)()
	}

	Exec(os.Stdout, cmdstr)
	Commands(cmdstr)()
}
