package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func getFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatal("Error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}

	return s
}

func getFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		log.Fatal("Error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}

	return i
}

func getFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		log.Fatal("Error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}

	return b
}
