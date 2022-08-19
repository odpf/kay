package cli

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/salt/cmdx"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "shield <command> <subcommand> [flags]",
		Short: "Kay is all-in-one tool for managing Apache Kafka clusters with confidence",
		Long: heredoc.Doc(`
			Kay is all-in-one tool for managing Apache Kafka clusters with confidence.
		`),
		SilenceUsage:  true,
		SilenceErrors: true,
		Example: heredoc.Doc(`
			$ kay cluster list
			$ kay topic list 

		`),
		Annotations: map[string]string{
			"group:core": "true",
			"help:learn": heredoc.Doc(`
				Use 'kay <command> <subcommand> --help' for more information about a command.
				Read the manual at https://odpf.github.io/kay/
			`),
			"help:feedback": heredoc.Doc(`
				Open an issue here https://github.com/odpf/kay/issues
			`),
			"help:environment": heredoc.Doc(`
				See 'kay help environment' for the list of supported environment variables.
			`),
		},
	}

	cmd.AddCommand(ServerCommand())

	// Help topics
	cmdx.SetHelp(cmd)
	cmd.AddCommand(cmdx.SetRefCmd(cmd))
	return cmd
}
