package cmd

import (
	"github.com/spf13/cobra"

	"github.com/techplexengineer/luabyexample/tools"
)

func RegisterServe(parentCmd *cobra.Command) *cobra.Command {

	var (
		// port to listen on
		port string
	)

	// serveCmd represents the serve command
	var serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "start a web server to view the generated site",
		Long: `The web server will rebuild the requested page 
on load to ensure up to date results are displayed.

Note the static server only listens on localhost and should not 
be used for production or network traffic.`,
		Run: func(cmd *cobra.Command, args []string) {
			tools.ServeRebuild(port)
		},
	}

	parentCmd.AddCommand(serveCmd)

	// allow the caller to change the directory and port
	parentCmd.Flags().StringVarP(&port, "port", "p", "8000", "Port to listen on")

	return serveCmd
}
