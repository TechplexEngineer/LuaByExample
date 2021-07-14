package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/techplexengineer/luabyexample/tools"
)

func RegisterTest(parentCmd *cobra.Command) *cobra.Command {
	// testCmd represents the test command
	testCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "test",
		Short: "Ensure lines are not too long to fit nicely in the layout",
		Long: `The html templates are designed for a fixed width, 
if code is too long it will spill out of the syntax highlighted area`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("Starting Measurement")
			tools.Measure()
			log.Println("Measurement Complete. All files pass.")
		},
	}

	parentCmd.AddCommand(testCmd)

	return testCmd
}
