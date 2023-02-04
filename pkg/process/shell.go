package process

import (
	"fmt"
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

	cmd, path, _ := strings.Cut(cmdstr, " ")
	path = strings.TrimSpace(path)
	switch cmd {
	case "cd":
		if path != "" {
			return func() error {
				return Chdir(path)
			}
		}

		return nop

	case "exist":
		if path != "" && PathNotExists(path) {
			return func() error { return fmt.Errorf("directory `%s' not exists", path) }
		}

		return nop

	case "not_exist":
		if path != "" && PathExists(path) {
			return func() error { return fmt.Errorf("directory `%s'  exists", path) }
		}

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
