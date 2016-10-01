package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type Selector struct {
	Value string
	Default string
	Options []string
}

func SelectorVar(cmd *cobra.Command, optionString string, defaultValue string, usage string, optionList []string) {
	cmd.Flags().Var(newSelector(defaultValue, optionList), optionString, usage)
}
	

func newSelector(defaultValue string, optionList []string) *Selector {
	return &Selector{
		Default: defaultValue,
		Options: optionList,
		Value: defaultValue,
	}
}

func (s *Selector) Set(inString string) error {
	if inString == "" {
		s.Value = s.Default
		return nil
	}

	// Normalize input string
	ts := strings.ToLower(inString)
	for _, o := range s.Options {
		if ts == o {
			s.Value = o
			return nil
		}
	}

	return fmt.Errorf("Invalid option: '%s'. Valid options are '%v'", inString, s.Options)
}

func (s *Selector) String() string { return s.Value }
func (s *Selector) Type() string { return "string" }
