package cmd

import (
	"github.com/spf13/cobra"
)

const defaultRootDir = "./public"

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "luabyexample",
		Short: "Toolbelt for developing/testing/building the site",
		Long: `LuaByExample is inspired by GoByExample and 
aims to be a gentle introduction to Lua with annotated examples.

This CLI tool transforms the examples into webpages that can be 
hosted by a static content system such as S3 or Github Pages.

The production site is available at https://luabyexample.techplexlabs.com/`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	RegisterBuild(rootCmd)
	RegisterClean(rootCmd)
	serveCmd := RegisterServe(rootCmd)
	RegisterServeStatic(serveCmd)
	RegisterTest(rootCmd)

	cobra.CheckErr(rootCmd.Execute())
}
