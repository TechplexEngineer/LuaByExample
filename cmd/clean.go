package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

func RegisterClean(parentCmd *cobra.Command) *cobra.Command {
	// cleanCmd represents the clean command
	cleanCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "clean",
		Short: "Remove built artifacts from the default directory",
		Long:  `Remove built artifacts from the default directory`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Print("Starting Clean")
			err := os.RemoveAll(defaultRootDir)
			if err != nil {
				log.Printf("error when removing %s directory - %s", defaultRootDir, err)
			}
			log.Print("Clean Complete")
		},
	}

	parentCmd.AddCommand(cleanCmd)

	return cleanCmd
}
