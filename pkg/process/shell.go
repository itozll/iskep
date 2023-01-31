package process

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

var nop = func() error { return nil }

func Command(cmdstr string) func() error {
	return CommandOutput(cmdstr, nil)
}

func Commands(cmdstrs []string) func() error {
	return func() error {
		return Exec(nil, cmdstrs)
	}
}

func CommandOutput(cmdstr string, out io.Writer) func() error {
	cmdstr = strings.TrimSpace(cmdstr)
	if len(cmdstr) == 0 {
		return nop
	}

	return func() error {
		cmd := exec.Command("/bin/bash", "-c", cmdstr)
		cmd.Stderr = os.Stderr
		cmd.Stdout = out
		return cmd.Run()
	}
}

func Exec(out io.Writer, cmdstrs []string) error {
	for _, cmdstr := range cmdstrs {
		err := CommandOutput(cmdstr, out)()
		if err != nil {
			return err
		}
	}

	return nil
}
