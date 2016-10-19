package commands

import (
	"github.com/spf13/cobra"
)

func newBranchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "branch",
		Short: "Repository Service Branches sub commands",
		Long: "Repository Service Branches sub commands",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.HelpFunc()(cmd, []string{})
			return nil
		},
	}

	cmd.AddCommand(newBranchGetCommand())
	cmd.AddCommand(newBranchListCommand())

	return cmd
}

func newBranchGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
		Short: "Get the specified branch for a repository",
		Long: "Get the specified branch for a repository",
		RunE: branchGet,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	cmd.Flags().String("branch", "", "Branch name")

	return cmd
}
func newBranchListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "List branches for the specified repository",
		Long: "List branches for the specified repository",
		RunE: branchList,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

func branchGet(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "branch", "owner", "repo"); err != nil {
		return err
	}

	branch, _, err := client.Repositories.GetBranch(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		getFlagString(cmd, "branch"),
		)
	if err != nil {
		return err
	}

	return output(cmd, branch)
}

func branchList(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	branches, _, err := client.Repositories.ListBranches(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil,
		)
	if err != nil {
		return err
	}

	return output(cmd, branches)
}

