/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/itozll/iskep/cmd/options"
	"github.com/itozll/iskep/pkg/runtime/iflag"
)

var appName = filepath.Base(os.Args[0])

// rootCmd represents the base command when called without any subcommands
var rootCmd = (&iflag.Command{
	Use: "iskep",
	PersistentArguments: []iflag.Argument{
		options.Verbose,
		options.DryRun,
		options.Dump,
	},
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Iskep is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Iskep application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}).Cobra()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
