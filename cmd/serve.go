package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/techplexengineer/luabyexample/tools"
	"log"
	"net/http"
	"os"
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

			//staticFiles := map[string]struct{}{
			//	"index.html": {},
			//	"monokai.css": {},
			//	"play.png": {},
			//	"site.css": {},
			//	"site.js": {},
			//	"favicon.ico": {},
			//	"clipboard.png": {},
			//}

			dir := "./static"
			dirEntries, err := os.ReadDir(dir)
			if err != nil {
				log.Print(fmt.Errorf("unable to list files in %s - %w", dir, err))
			}
			for _, entry := range dirEntries {
				if entry.IsDir() {
					continue
				}
				fileName := entry.Name()
				http.HandleFunc("/"+fileName, func(w http.ResponseWriter, r *http.Request) {

					log.Printf("Request for %s", "/"+fileName)
					http.ServeFile(w, r, "./static/"+fileName)
				})
			}
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				//examples := tools.ParseExamples()
				tools.RenderIndex(w, tools.GetExampleNamesAndIds("./examples.txt"))
				//if _, ok := staticFiles[strings.Trim(r.RequestURI, "/")]; ok {
				//
				//}

				//_, _ = fmt.Fprintf(w, "test \n %s", r.RequestURI)

				//_, _ = w.Write([]byte(fmt.print))
			})

			address := "127.0.0.1:" + port
			fmt.Printf("Starting server on http://%s\n", address)
			err = http.ListenAndServe(address, nil)
			if err != nil {
				panic(err)
			}

		},
	}

	parentCmd.AddCommand(serveCmd)

	// allow the caller to change the directory and port
	parentCmd.Flags().StringVarP(&port, "port", "p", "8000", "Port to listen on")

	return serveCmd
}
