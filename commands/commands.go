package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init(name, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use: name,
		Short: "Command line interface for Github REST API",
		Long: "Command line interface for Github REST API",
		SilenceUsage: true,
		SilenceErrors: true,
		RunE: func(c *cobra.Command, args []string) error {
			c.HelpFunc()(c, []string{})
			return nil
		},
	}

	addPersistentFlags(cmd)
	cmd.AddCommand(newVersionCommand(name, version))
	cmd.AddCommand(newRepositoryCommand())

	return cmd
}

func newVersionCommand(name, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "version",
		Short: "Print version information",
		Long: "Print version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s %s\n", name, version)
			return nil
		},
	}

	return cmd
}
