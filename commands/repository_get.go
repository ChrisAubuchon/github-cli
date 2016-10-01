package commands

import (
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func newRepositoryGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
		Short: "Get fetches a repository",
		Long: "Get fetches a repository",
		RunE: repositoryGet,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	cmd.AddCommand(newRepositoryGetRelease())

	return cmd
}

func newRepositoryGetRelease() *cobra.Command {
	cmd := &cobra.Command{
		Use: "release",
		Short: "Fetch a single release",
		Long: "Fetch a single release",
		RunE: repositoryGetRelease,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addIdIntFlag(cmd, "Release id. Default is latest release")
	addTemplateFlags(cmd)

	return cmd
}

func repositoryGet(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	repo, _, err := client.Repositories.Get(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		)
	if err != nil {
		return err
	}

	return output(cmd, repo)
}

func repositoryGetRelease(cmd *cobra.Command, args []string) error {
	var err error
	var release *github.RepositoryRelease

	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	id := getFlagInt(cmd, "id")
	if id == 0 {
		release, _, err = client.Repositories.GetLatestRelease(
			getFlagString(cmd, "owner"),
			getFlagString(cmd, "repo"),
			)
	} else {
		release, _, err = client.Repositories.GetRelease(
			getFlagString(cmd, "owner"),
			getFlagString(cmd, "repo"),
			id,
			)
	}

	if err != nil {
		return err
	}

	return output(cmd, release)
}
