package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func output(cmd *cobra.Command, v interface{}) error {
	tflag := getFlagString(cmd, "template")

	if tflag == "" {
		return outputJson(v)
	} else {
		return outputTemplate(tflag, v)
	}
}

func outputJson(v interface{}) error {
	jsonRaw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	jsonStr := string(jsonRaw)
	if strings.HasSuffix(jsonStr, "\n") {
		fmt.Printf(jsonStr)
	} else {
		fmt.Printf(jsonStr + "\n")
	}

	return nil
}

func outputTemplate(t string, v interface{}) error {
	if strings.HasPrefix(t, "@") {
		v, err := ioutil.ReadFile(t[1:])
		if err != nil {
			return err
		}
		t = string(v)
	}

	tparsed, err := template.New("").Parse(t)
	if err != nil {
		return err
	}

	return tparsed.Execute(os.Stdout, v)
}
