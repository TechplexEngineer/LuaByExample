package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func RegisterClean(parentCmd *cobra.Command) *cobra.Command {

	// cleanCmd represents the clean command
	var cleanCmd = &cobra.Command{
		Use:   "clean",
		Short: "Remove built artifacts from the default directory",
		Long:  `Remove built artifacts from the default directory`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("clean called")
			err := os.RemoveAll(defaultRootDir)
			if err != nil {
				log.Printf("error when removing %s directory - %s", defaultRootDir, err)
			}
		},
	}

	parentCmd.AddCommand(cleanCmd)

	return cleanCmd
}
