package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/symonk/rdsmon-go/version"
)

// versionCommand represents the version command
var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Prints the rdsmon-go version and exits.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Date:", version.BuildDate)
		fmt.Println("Git Commit:", version.GitCommit)
		fmt.Println("Version:", version.Version)
		fmt.Println("Go Version:", version.GoVersion)
		fmt.Println("OS / Arch:", version.OsArch)
	},
}

func init() {
	rootCommand.AddCommand(versionCommand)
}
