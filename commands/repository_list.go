package commands

import (
	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func newRepositoryListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Short: "List the repositories for a user",
		Long: "List the repositories for a user",
		RunE: repositoryList,
	}

	cmd.Flags().String("user", "", "Github user")
	SelectorVar(
		cmd, 
		"visibility", 
		"", 
		"Visibility of repositories to list. One of [all|public|private]",
		[]string{ "all", "public", "private" },
		)
	SelectorVar(
		cmd,
		"type",
		"",
		"Type of repositories to list. One of [all|owner|public|private|member]",
		[]string{ "all", "owner", "public", "private", "member" },
		)
	SelectorVar(
		cmd,
		"sort",
		"",
		"How to sort the repository list. One of [created|updated|pushed|full_name]",
		[]string{ "created", "updated", "pushed", "full_name" },
		)
	SelectorVar(
		cmd,
		"sort-direction",
		"",
		"Direction in which to sort repositories. One of [asc|desc]",
		[]string{ "asc", "desc" },
		)
	ListVar(
		cmd,
		"affiliation",
		"",
		"List repos of a given affillation. Comma separated list of values. Can include owner, collaborator, organization_member",
		[]string{ "owner", "collaborator", "organization_member" },
		)
	addTemplateFlags(cmd)

	cmd.AddCommand(newRepositoryListBranches())
	cmd.AddCommand(newRepositoryListReleases())

	return cmd
}

func newRepositoryListBranches() *cobra.Command {
	cmd := &cobra.Command{
		Use: "branches",
		Short: "List the branches for a repository",
		Long: "List the branches for a repository",
		RunE: repositoryListBranches,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

func newRepositoryListReleases() *cobra.Command {
	cmd := &cobra.Command{
		Use: "releases",
		Short: "List the releases for a repository",
		Long: "List the releases for a repository",
		RunE: repositoryListReleases,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

func repositoryList(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	repos, _, err := client.Repositories.List(getFlagString(cmd, "user"), &github.RepositoryListOptions{
		Visibility: getFlagString(cmd, "visibility"),
		Affiliation: getFlagString(cmd, "affiliation"),
		Type: getFlagString(cmd, "type"),
		Sort: getFlagString(cmd, "sort"),
		Direction: getFlagString(cmd, "sort-direction"),
		})
	if err != nil {
		return err
	}

	return output(cmd, repos)
}

func repositoryListBranches(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	branches, _, err := client.Repositories.ListBranches(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil)
	if err != nil {
		return err
	}

	return output(cmd, branches)
}

func repositoryListReleases(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	release, _, err := client.Repositories.ListReleases(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil)
	if err != nil {
		return err
	}

	return output(cmd, release)
}
