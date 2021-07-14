package tools

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ServeStatic(port string, rootDir string) {
	err := http.ListenAndServe("127.0.0.1:"+port, http.FileServer(http.Dir(rootDir)))
	if err != nil {
		panic(err)
	}
}

func ServeRebuild(port string) {
	r := mux.NewRouter()
	r.StrictSlash(true)

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

		r.HandleFunc("/"+fileName, func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request for %s", "/"+fileName)
			http.ServeFile(w, r, "./static/"+fileName)
		})
	}

	r.HandleFunc("/{exampleId}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		example := ParseExample(vars["exampleId"]) // @todo really need example name not id
		example.PrevExample, example.NextExample = GetPrevNextExample(vars["exampleId"])

		RenderExample(w, example)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		examples := GetExampleNamesAndIds("./examples.txt")
		RenderIndex(w, examples)
	})

	address := "127.0.0.1:" + port
	log.Printf("Starting server on http://%s\n", address)
	check(http.ListenAndServe(address, r))
}
