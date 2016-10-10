package commands

import (
	"github.com/google/go-github/github"
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

	cmd.AddCommand(newRepositoryListByOrg())
	cmd.AddCommand(newRepositoryListContributors())
	cmd.AddCommand(newRepositoryListLanguages())
	cmd.AddCommand(newRepositoryListTags())
	cmd.AddCommand(newRepositoryListTeams())
//	cmd.AddCommand(newRepositoryListBranches())
//	cmd.AddCommand(newRepositoryListCollaborators())
//	cmd.AddCommand(newRepositoryListComments())
//	cmd.AddCommand(newRepositoryListReleases())

	return cmd
}

func newRepositoryListByOrg() *cobra.Command {
	cmd := &cobra.Command{
		Use: "byorg",
		Short: "List the repositories for an organization",
		Long: "List the repositories for an organization",
		RunE: repositoryListByOrg,
	}

	SelectorVar(
		cmd,
		"type",
		"",
		"Type of repositories to list. One of [all|public|private|forks|sources|member]",
		[]string{ "all", "public", "private", "forks", "sources", "member" },
		)

	cmd.Flags().String("org", "", "Organization name")

	addTemplateFlags(cmd)

	return cmd
}
	

func newRepositoryListContributors() *cobra.Command {
	cmd := &cobra.Command{
		Use: "contributors",
		Short: "List contributors for a repository",
		Long: "List contributors for a repository",
		RunE: repositoryListContributors,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	cmd.Flags().Bool("anon", false, "Include anonymous contributors in results")

	return cmd
}

func newRepositoryListLanguages() *cobra.Command {
	cmd := &cobra.Command{
		Use: "languages",
		Short: "List languages for the specified repository",
		Long: "List languages for the specified repository",
		RunE: repositoryListLanguages,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

func newRepositoryListTags() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tags",
		Short: "List tags for the specified repository",
		Long: "List tags for the specified repository",
		RunE: repositoryListTags,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

func newRepositoryListTeams() *cobra.Command {
	cmd := &cobra.Command{
		Use: "teams",
		Short: "List the teams for the specified repository",
		Long: "List the teamss for the specified repository",
		RunE: repositoryListTeams,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	return cmd
}

//func newRepositoryListBranches() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "branches",
//		Short: "List the branches for a repository",
//		Long: "List the branches for a repository",
//		RunE: repositoryListBranches,
//	}
//
//	addOwnerFlag(cmd)
//	addRepoFlag(cmd)
//	addTemplateFlags(cmd)
//
//	return cmd
//}
//
//func newRepositoryListCollaborators() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "collaborators",
//		Short: "List the Github users that have access to a repository",
//		Long: "List the Github users that have access to a repository",
//		RunE: repositoryListCollaborators,
//	}
//
//	addOwnerFlag(cmd)
//	addRepoFlag(cmd)
//	addTemplateFlags(cmd)
//
//	return cmd
//}

//func newRepositoryListComments() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "comments",
//		Short: "List all of the comments for a repository",
//		Long: "List all of the comments for a repository",
//		RunE: repositoryListComments,
//	}
//
//	addOwnerFlag(cmd)
//	addRepoFlag(cmd)
//	addTemplateFlags(cmd)
//
//	return cmd
//}
//
//func newRepositoryListReleases() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "releases",
//		Short: "List the releases for a repository",
//		Long: "List the releases for a repository",
//		RunE: repositoryListReleases,
//	}
//
//	addOwnerFlag(cmd)
//	addRepoFlag(cmd)
//	addTemplateFlags(cmd)
//
//	return cmd
//}

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

func repositoryListByOrg(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "org"); err != nil {
		return err
	}

	repos, _, err := client.Repositories.ListByOrg(
		getFlagString(cmd, "org"),
		&github.RepositoryListByOrgOptions{
			Type: getFlagString(cmd, "type"),
		},
	)
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

func repositoryListCollaborators(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	collaborators, _, err := client.Repositories.ListCollaborators(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil)
	if err != nil {
		return err
	}

	return output(cmd, collaborators)
}

func repositoryListContributors(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	var lco *github.ListContributorsOptions
	if getFlagBool(cmd, "anon") {
		lco = &github.ListContributorsOptions{
			Anon: "true",
		}
	} else {
		lco = nil
	}

	contributors, _, err := client.Repositories.ListContributors(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		lco,
		)
	if err != nil {
		return err
	}

	return output(cmd, contributors)
}

func repositoryListLanguages(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	languages, _, err := client.Repositories.ListLanguages(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		)
	if err != nil {
		return err
	}

	return output(cmd, languages)
}

func repositoryListTags(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	tags, _, err := client.Repositories.ListTags(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil,
		)
	if err != nil {
		return err
	}

	return output(cmd, tags)
}

func repositoryListTeams(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	teams, _, err := client.Repositories.ListTeams(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil,
		)
	if err != nil {
		return err
	}

	return output(cmd, teams)
}

func repositoryListComments(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {	
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	comments, _, err := client.Repositories.ListComments(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		nil)
	if err != nil {
		return err
	}

	return output(cmd, comments)
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

	return cmd
}

//func newRepositoryGetRelease() *cobra.Command {
//	cmd := &cobra.Command{
//		Use: "release",
//		Short: "Fetch a single release",
//		Long: "Fetch a single release",
//		RunE: repositoryGetRelease,
//	}
//
//	addOwnerFlag(cmd)
//	addRepoFlag(cmd)
//	addIdIntFlag(cmd, "Release id. Default is latest release")
//	addTemplateFlags(cmd)
//
//	return cmd
//}

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
