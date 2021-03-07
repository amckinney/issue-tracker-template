package command

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Command is a command.
type Command struct {
	// Use is the usage message.
	Use string
	// Short is the short message shown in the 'help' output.
	Short string
	// Long is the long message shown in the 'help' output.
	Long string
	// Args are the expected arguments.
	Args cobra.PositionalArgs
	// BindFlags allows binding of flags on build.
	BindFlags func(*pflag.FlagSet)
	// Run is the command to run.
	Run func(context.Context) error
	// Version is the version of this command.
	Version string
}
