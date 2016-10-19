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

	cmd.AddCommand(newRepositoryCreateCommand())
	cmd.AddCommand(newRepositoryDeleteCommand())
	cmd.AddCommand(newRepositoryEditCommand())
	cmd.AddCommand(newRepositoryGetCommand())
	cmd.AddCommand(newRepositoryListCommand())

	return cmd
}

func newRepositoryCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
		Short: 	"Create a new repository",
		Long: "Create a new repository",
		RunE: repositoryCreate,
	}

	cmd.Flags().Bool("auto-init", false, "Create initial commit with empty README")
	cmd.Flags().String("description", "", "A short description of the repository")
	cmd.Flags().String("gitignore-template", "", "Desired language or platform .gitignore template to apply")
	cmd.Flags().String("homepage", "", "A URL with more information about the repository")
	cmd.Flags().String("license-template", "", "Desired license template to apply")
	cmd.Flags().String("name", "", "Repository name")
	cmd.Flags().Bool("no-downloads", false, "Disable downloads for the repository")
	cmd.Flags().Bool("no-issues", false, "Disable issues on the repository")
	cmd.Flags().Bool("no-wiki", false, "Disable the wiki for the repository")
	cmd.Flags().String("org", "", "Organization name")
	cmd.Flags().Bool("private", false, "Create a private repository")
	cmd.Flags().Int("team-id", 0, "Id of the team that will be granted access to this repository. Only valid with --org")

	addTemplateFlags(cmd)

	return cmd
}

func newRepositoryDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Short: 	"Delete a repository",
		Long: "Delete a repository",
		RunE: repositoryDelete,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)

	return cmd
}


func newRepositoryEditCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit",
		Short: 	"Update a repository",
		Long: "Update a repository",
		RunE: repositoryEdit,
	}

	addOwnerFlag(cmd)
	addRepoFlag(cmd)
	addTemplateFlags(cmd)

	cmd.Flags().String("description", "", "A short description of the repository")
	cmd.Flags().String("homepage", "", "A URL with more information about the repository")
	cmd.Flags().String("name", "", "Repository name")
	cmd.Flags().Bool("no-downloads", false, "Disable downloads for the repository")
	cmd.Flags().Bool("no-issues", false, "Disable issues on the repository")
	cmd.Flags().Bool("no-wiki", false, "Disable the wiki for the repository")
	cmd.Flags().Bool("private", false, "Create a private repository")
	cmd.Flags().String("default-branch", "", "Updates the default branch for the repository")

	return cmd
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

func repositoryCreate(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "name"); err != nil {
		return err
	}

	r := new(github.Repository)
	if getFlagBool(cmd, "auto-init") {
		r.AutoInit = new(bool)
		*r.AutoInit = true
	}

	if desc := getFlagString(cmd, "description"); desc != "" {
		r.Description = new(string)
		*r.Description = desc
	}

	if gt := getFlagString(cmd, "gitignore-template"); gt != "" {
		r.GitignoreTemplate = new(string)
		*r.GitignoreTemplate = gt
	}

	if h := getFlagString(cmd, "homepage"); h != "" {
		r.Homepage = new(string)
		*r.Homepage = h
	}

	if getFlagBool(cmd, "no-downloads") {
		r.HasDownloads = new(bool)
		*r.HasDownloads = false
	}

	if getFlagBool(cmd, "no-issues") {
		r.HasIssues = new(bool)
		*r.HasIssues = false
	}

	if getFlagBool(cmd, "no-wiki") {
		r.HasWiki = new(bool)
		*r.HasWiki = false
	}

	if h := getFlagString(cmd, "homepage"); h != "" {
		r.Homepage = new(string)
		*r.Homepage = h
	}

	if lt := getFlagString(cmd, "license-template"); lt != "" {
		r.LicenseTemplate = new(string)
		*r.LicenseTemplate = lt
	}

	if getFlagBool(cmd, "private") {
		r.Private = new(bool)
		*r.Private = true
	}

	if i := getFlagInt(cmd, "team-id"); i != 0 {
		r.TeamID = new(int)
		*r.TeamID = i
	}

	r.Name = new(string)
	*r.Name = getFlagString(cmd, "name")

	repo, _, err := client.Repositories.Create(getFlagString(cmd, "org"), r)
	if err != nil {
		return err
	}

	return output(cmd, repo)
}

func repositoryDelete(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo"); err != nil {
		return err
	}

	if _, err := client.Repositories.Delete(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
	); err != nil {
		return err
	}

	return nil
}

func repositoryEdit(cmd *cobra.Command, args []string) error {
	client, err := newClient(cmd)
	if err != nil {
		return err
	}

	if err := validateFlagsNotEmpty(cmd, "owner", "repo", "name"); err != nil {
		return err
	}

	r := new(github.Repository)
	if d := getFlagString(cmd, "default-branch"); d != "" {
		r.DefaultBranch = new(string)
		*r.DefaultBranch = d
	}

	if desc := getFlagString(cmd, "description"); desc != "" {
		r.Description = new(string)
		*r.Description = desc
	}

	if h := getFlagString(cmd, "homepage"); h != "" {
		r.Homepage = new(string)
		*r.Homepage = h
	}

	if getFlagBool(cmd, "no-downloads") {
		r.HasDownloads = new(bool)
		*r.HasDownloads = false
	}

	if getFlagBool(cmd, "no-issues") {
		r.HasIssues = new(bool)
		*r.HasIssues = false
	}

	if getFlagBool(cmd, "no-wiki") {
		r.HasWiki = new(bool)
		*r.HasWiki = false
	}

	if getFlagBool(cmd, "private") {
		r.Private = new(bool)
		*r.Private = true
	}

	r.Name = new(string)
	*r.Name = getFlagString(cmd, "name")

	repo, _, err := client.Repositories.Edit(
		getFlagString(cmd, "owner"),
		getFlagString(cmd, "repo"),
		r,
	)
	if err != nil {
		return err
	}

	return output(cmd, repo)
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
