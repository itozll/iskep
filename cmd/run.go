/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/itozll/iskep/cmd/options"
	"github.com/itozll/iskep/pkg/runtime/iflag"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &iflag.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Arguments: []iflag.Argument{
		options.GroupName,
	},
	Long: `Add (iskep run) will create a run command, with a license and
the appropriate structure for a Cobra-based CLI application,
and register it to its parent (default rootCmd).

If you want your command to be public, pass in the command name
with an initial uppercase letter.

Example: iskep run server -> resulting in a run cmd/server.go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called", options.GroupName.Value())
		fmt.Printf("%+v\n", options.GoVersion.Value())
		fmt.Printf("%+v\n", options.Verbose.Value())
		fmt.Println(options.GroupName.Value())
	},
}

func init() {
	rootCmd.AddCommand(runCmd.Cobra())
}
