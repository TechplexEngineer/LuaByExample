package tools

import (
	"net/http"
)

func ServeStatic(port string, rootDir string) {
	err := http.ListenAndServe("127.0.0.1:"+port, http.FileServer(http.Dir(rootDir)))
	if err != nil {
		panic(err)
	}
}
