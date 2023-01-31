/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/itozll/iskep/cmd/options"
	"github.com/itozll/iskep/internal/etcd"
	"github.com/itozll/iskep/pkg/process/build"
	"github.com/itozll/iskep/pkg/runtime/iflag"
	"github.com/itozll/iskep/pkg/runtime/rtinfo"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &iflag.Command{
	Use:     "new [flags] <workspace>",
	Aliases: []string{"n"},
	Short:   "create an go workspace",
	Arguments: []iflag.Argument{
		options.GoVersion,
		options.GroupName,
		options.Path,
		options.Local,
	},
	Example: fmt.Sprintf(`  %s new --group mygroup myrepos
  %s new mygroup/myrepos
`, appName, appName),
	SilenceUsage: true,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Help()
			os.Exit(1)
		}

		return rtinfo.Init(args[0])
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		command, err := build.NewCommand(etcd.NewCommandConfig)
		if err != nil {
			return err
		}

		switch {
		case options.Local.Value():
			rtinfo.TargetPath = ""
		case rtinfo.TargetPath == "":
			rtinfo.TargetPath = rtinfo.Info.Project
		case rtinfo.TargetPath != "":
			rtinfo.TargetPath = strings.TrimRight(rtinfo.TargetPath, "/") + "/" + rtinfo.Info.Project
		}

		command.AttachMap(rtinfo.Binder())
		command.Attach("command", "server")
		command.Attach("parent_cmd", "root")
		return command.Exec(nil)
	},
}

func init() {
	rootCmd.AddCommand(newCmd.Cobra())
}
