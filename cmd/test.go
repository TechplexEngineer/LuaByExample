package cmd

import (
	"fmt"
	"github.com/techplexengineer/luabyexample/tools"

	"github.com/spf13/cobra"
)

func RegisterTest(parentCmd *cobra.Command) *cobra.Command {

	// testCmd represents the test command
	var testCmd = &cobra.Command{
		Use:   "test",
		Short: "Ensure lines are not too long to fit nicely in the layout",
		Long: `The html templates are designed for a fixed width, 
if code is too long it will spill out of the syntax highlighted area`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Measurement")
			tools.Measure()
			fmt.Println("Measurement Complete. All files pass.")
		},
	}

	parentCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return testCmd
}
