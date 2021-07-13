package tools

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func ServeStatic(port string, rootDir string) {
	err := http.ListenAndServe("127.0.0.1:"+port, http.FileServer(http.Dir(rootDir)))
	if err != nil {
		panic(err)
	}
}

func ServeRebuild(port string) {
	r := mux.NewRouter()

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

		example := ParseExample(vars["exampleId"]) //@todo really need example name not id
		example.PrevExample, example.NextExample = GetPrevNextExample(vars["exampleId"])

		RenderExample(w, example)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		examples := GetExampleNamesAndIds("./examples.txt")
		RenderIndex(w, examples)
	})

	address := "127.0.0.1:" + port
	fmt.Printf("Starting server on http://%s\n", address)
	err = http.ListenAndServe(address, r)
	if err != nil {
		panic(err)
	}
}
