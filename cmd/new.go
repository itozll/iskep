/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/itozll/iskep/cmd/options"
	"github.com/itozll/iskep/runtime/iflag"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &iflag.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Arguments: []iflag.Argument{
		options.GroupName,
	},
	Long: `Add (iskep new) will create a new command, with a license and
the appropriate structure for a Cobra-based CLI application,
and register it to its parent (default rootCmd).

If you want your command to be public, pass in the command name
with an initial uppercase letter.

Example: iskep new server -> resulting in a new cmd/server.go`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called", options.GroupName.Value())
		fmt.Printf("%+v\n", options.GoVersion.Value())
		fmt.Printf("%+v\n", options.Verbose.Value())
		fmt.Println(options.GroupName.Value())
	},
}

func init() {
	rootCmd.AddCommand(newCmd.Cobra())
}
