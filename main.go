package main

import (
	"fmt"
	"os"

	"github.com/ChrisAubuchon/github-cli/commands"
)

const Name = "github-cli"
const Version = "0.1.0"

func main() {
	root := commands.Init(Name, Version)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
