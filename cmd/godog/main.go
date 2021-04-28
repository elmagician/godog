package main

import (
	"github.com/elmagician/godog/cmd/godog/internal"
)

func main() {
	rootCmd := internal.CreateRootCmd()
	buildCmd := internal.CreateBuildCmd()
	runCmd := internal.CreateRunCmd()
	versionCmd := internal.CreateVersionCmd()

	rootCmd.AddCommand(&buildCmd, &runCmd, &versionCmd)
	rootCmd.Execute()
}
