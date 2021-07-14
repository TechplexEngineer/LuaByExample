package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/techplexengineer/luabyexample/tools"
)

func RegisterServeStatic(parentCmd *cobra.Command) *cobra.Command {
	var (
		// port to listen on
		port string
		// root directory for server
		rootDir string
	)

	// serveStaticCmd represents the serveStatic command
	serveStaticCmd := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "static",
		Short: "Start http file server for an existing built site. Does not rebuild.",
		Long: `It can be helpful to start a static file server to test changes to the build process.

Note the static server only listens on localhost and should not 
be used for production or network traffic.
`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("Starting static server on %s for dir %s\n", port, rootDir)

			tools.ServeStatic(port, rootDir)
		},
	}

	// static is a sub-command of serve
	parentCmd.AddCommand(serveStaticCmd)

	// allow the caller to change the directory and port
	parentCmd.Flags().StringVarP(&port, "port", "p", "8000", "Port to listen on")
	parentCmd.Flags().StringVarP(&rootDir, "directory", "d", defaultRootDir, "directory to serve")

	return serveStaticCmd
}
