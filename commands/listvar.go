package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type listVar struct {
	Value string
	Default string
	Options []string
}

func ListVar(cmd *cobra.Command, optionString string, defaultValue string, usage string, optionList []string) {
	cmd.Flags().Var(newList(defaultValue, optionList), optionString, usage)
}
	

func newList(defaultValue string, optionList []string) *listVar {
	return &listVar{
		Default: defaultValue,
		Options: optionList,
		Value: defaultValue,
	}
}

func (s *listVar) Set(inString string) error {
	if inString == "" {
		s.Value = s.Default
		return nil
	}

	// Normalize input string
	ts := strings.ToLower(inString)
	for _, t := range strings.Split(ts, ",") {
		if !stringInSlice(t, s.Options) {
			return fmt.Errorf("Invalid option: '%s'. Valid options are '%v'", t, s.Options)
		}
	}
	s.Value = ts

	return nil
}

func (s *listVar) String() string { return s.Value }
func (s *listVar) Type() string { return "string" }
