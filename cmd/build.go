package cmd

import (
	"fmt"
	"github.com/techplexengineer/luabyexample/tools"

	"github.com/spf13/cobra"
)

func RegisterBuild(parentCmd *cobra.Command) *cobra.Command {

	var siteDir string

	// buildCmd represents the build command
	var buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Generate static web content from examples.",
		Long:  `Parse code examples, execute shell files to capture output, render templates to public directory`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("build called")
			tools.Generate(siteDir)
		},
	}

	parentCmd.AddCommand(buildCmd)

	// directory flag applies only to build command
	buildCmd.Flags().StringVarP(&siteDir, "directory", "d", defaultRootDir, "Directory to output built site into")
	return buildCmd
}
