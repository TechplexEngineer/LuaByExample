package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/techplexengineer/luabyexample/tools"
)

func RegisterBuild(parentCmd *cobra.Command) *cobra.Command {
	var siteDir string

	// buildCmd represents the build command
	buildCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "build",
		Short: "Generate static web content from examples.",
		Long:  `Parse code examples, execute shell files to capture output, render templates to public directory`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Starting Build")
			tools.Generate(siteDir)
			log.Print("Build Completed")
		},
	}

	parentCmd.AddCommand(buildCmd)

	// directory flag applies only to build command
	buildCmd.Flags().StringVarP(&siteDir, "directory", "d", defaultRootDir, "Directory to output built site into")

	return buildCmd
}
