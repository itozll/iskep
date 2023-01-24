package iflag

import (
	"github.com/spf13/cobra"
)

type Command struct {
	// Use is the one-line usage message.
	// Recommended syntax is as follow:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string

	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string

	// arguments
	Arguments []Argument

	// persistent arguments
	PersistentArguments []Argument

	// SuggestFor is an array of command names for which this command will be suggested -
	// similar to aliases but only suggests.
	SuggestFor []string

	// Short is the short description shown in the 'help' output.
	Short string

	// The group id under which this subcommand is grouped in the 'help' output of its parent.
	GroupID string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Example is examples of how to use the command.
	Example string

	// ValidArgs is list of all valid non-flag arguments that are accepted in shell completions
	ValidArgs []string
	// ValidArgsFunction is an optional function that provides valid non-flag arguments for shell completion.
	// It is a dynamic version of using ValidArgs.
	// Only one of ValidArgs and ValidArgsFunction can be used for a command.
	ValidArgsFunction func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective)

	// Expected arguments
	Args cobra.PositionalArgs

	// ArgAliases is List of aliases for ValidArgs.
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	ArgAliases []string

	// BashCompletionFunction is custom bash functions used by the legacy bash autocompletion generator.
	// For portability with other shells, it is recommended to instead use ValidArgsFunction
	BashCompletionFunction string

	// Deprecated defines, if this command is deprecated and should print this string when used.
	Deprecated string

	// Annotations are key/value pairs that can be used by applications to identify or
	// group commands.
	Annotations map[string]string

	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
	// command does not define one.
	Version string

	// The *Run functions are executed in the following order:
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// All functions get the same args, the arguments after the command name.
	//
	// PersistentPreRun: children of this command will inherit and execute.
	PersistentPreRun func(cmd *cobra.Command, args []string)
	// PersistentPreRunE: PersistentPreRun but returns an error.
	PersistentPreRunE func(cmd *cobra.Command, args []string) error
	// PreRun: children of this command will not inherit.
	PreRun func(cmd *cobra.Command, args []string)
	// PreRunE: PreRun but returns an error.
	PreRunE func(cmd *cobra.Command, args []string) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *cobra.Command, args []string)
	// RunE: Run but returns an error.
	RunE func(cmd *cobra.Command, args []string) error
	// PostRun: run after the Run command.
	PostRun func(cmd *cobra.Command, args []string)
	// PostRunE: PostRun but returns an error.
	PostRunE func(cmd *cobra.Command, args []string) error
	// PersistentPostRun: children of this command will inherit and execute after PostRun.
	PersistentPostRun func(cmd *cobra.Command, args []string)
	// PersistentPostRunE: PersistentPostRun but returns an error.
	PersistentPostRunE func(cmd *cobra.Command, args []string) error

	// FParseErrWhitelist flag parse errors to be ignored
	FParseErrWhitelist cobra.FParseErrWhitelist

	// CompletionOptions is a set of options to control the handling of shell completion
	CompletionOptions cobra.CompletionOptions

	// TraverseChildren parses flags on all parents before executing child command.
	TraverseChildren bool

	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
	Hidden bool

	// SilenceErrors is an option to quiet errors down stream.
	SilenceErrors bool

	// SilenceUsage is an option to silence usage when an error occurs.
	SilenceUsage bool

	// DisableFlagParsing disables the flag parsing.
	// If this is true all flags will be passed to the command as arguments.
	DisableFlagParsing bool

	// DisableAutoGenTag defines, if gen tag ("Auto generated by spf13/cobra...")
	// will be printed by generating docs for this command.
	DisableAutoGenTag bool

	// DisableFlagsInUseLine will disable the addition of [flags] to the usage
	// line of a command when printing help or generating docs
	DisableFlagsInUseLine bool

	// DisableSuggestions disables the suggestions based on Levenshtein distance
	// that go along with 'unknown command' messages.
	DisableSuggestions bool

	// SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
	// Must be > 0.
	SuggestionsMinimumDistance int
}

func (cmd *Command) Cobra() *cobra.Command {
	c := &cobra.Command{
		Use:                        cmd.Use,
		Aliases:                    cmd.Aliases,
		SuggestFor:                 cmd.SuggestFor,
		Short:                      cmd.Short,
		GroupID:                    cmd.GroupID,
		Long:                       cmd.Long,
		Example:                    cmd.Example,
		ValidArgs:                  cmd.ValidArgs,
		ValidArgsFunction:          cmd.ValidArgsFunction,
		Args:                       cmd.Args,
		ArgAliases:                 cmd.ArgAliases,
		BashCompletionFunction:     cmd.BashCompletionFunction,
		Deprecated:                 cmd.Deprecated,
		Annotations:                cmd.Annotations,
		Version:                    cmd.Version,
		PersistentPreRun:           cmd.PersistentPreRun,
		PersistentPreRunE:          cmd.PersistentPreRunE,
		PreRun:                     cmd.PreRun,
		PreRunE:                    cmd.PreRunE,
		Run:                        cmd.Run,
		RunE:                       cmd.RunE,
		PostRun:                    cmd.PostRun,
		PostRunE:                   cmd.PostRunE,
		PersistentPostRun:          cmd.PersistentPostRun,
		PersistentPostRunE:         cmd.PersistentPostRunE,
		FParseErrWhitelist:         cmd.FParseErrWhitelist,
		CompletionOptions:          cmd.CompletionOptions,
		TraverseChildren:           cmd.TraverseChildren,
		Hidden:                     cmd.Hidden,
		SilenceErrors:              cmd.SilenceErrors,
		SilenceUsage:               cmd.SilenceUsage,
		DisableFlagParsing:         cmd.DisableFlagParsing,
		DisableAutoGenTag:          cmd.DisableAutoGenTag,
		DisableFlagsInUseLine:      cmd.DisableFlagsInUseLine,
		DisableSuggestions:         cmd.DisableSuggestions,
		SuggestionsMinimumDistance: cmd.SuggestionsMinimumDistance,
	}

	for _, arg := range cmd.Arguments {
		arg.Bind(c.Flags())
	}

	for _, arg := range cmd.PersistentArguments {
		arg.Bind(c.PersistentFlags())
	}

	return c
}
