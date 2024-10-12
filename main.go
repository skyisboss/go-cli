package main

import "go-cli/cmd"

var (
	gitCommit  string
	gitVersion string
)

func main() {
	cmd.Version = gitVersion
	cmd.Commit = gitCommit
	cmd.Execute()
}
