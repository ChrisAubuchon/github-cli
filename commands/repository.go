package commands

import (
	"github.com/spf13/cobra"
)

func newRepositoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "repo",
		Short: "Repository Service sub commands",
		Long: "Repository Service sub commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, []string{})
			return nil
		},
	}

	cmd.AddCommand(newRepositoryGetCommand())
	cmd.AddCommand(newRepositoryListCommand())

	return cmd
}
