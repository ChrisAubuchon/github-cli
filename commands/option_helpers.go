package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func addPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String("token", "", "OAuth2 token")
	cmd.PersistentFlags().String("username", "", "Github user name")
	cmd.PersistentFlags().String("password", "", "Github password")
	cmd.PersistentFlags().String("otp", "", "Github one time password for two-factor authentication")
}

func addTemplateFlags(cmd *cobra.Command) {
	cmd.Flags().String("template", "", "Output template. Use @filename to read template from a file")
}

func addOwnerFlag(cmd *cobra.Command) {
	cmd.Flags().String("owner", "", "Repository owner")
}

func addRepoFlag(cmd *cobra.Command) {
	cmd.Flags().String("repo", "", "Repository name")
}

func addIdIntFlag(cmd *cobra.Command, desc string) {
	cmd.Flags().Int("id", 0, desc)
}

func validateFlagsNotEmpty(cmd *cobra.Command, opts ...string) error {
	for _, o := range opts {
		if getFlagString(cmd, o) == "" {
			return fmt.Errorf("Required option '%s' not set for command '%s'", o, cmd.Name())
		}
	}

	return nil
}
