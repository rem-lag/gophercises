package main

import (
	"fmt"
	"net/http"
	"shorturl/short"
)

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!!!")
}

func main() {

	mux := defaultMux()
	mapPaths := map[string]string{
		"/rem-lag":  "https://github.com/rem-lag",
		"/cv-stack": "https://stats.stackexchange.com",
	}

	mapHandler := short.MapHandler(mapPaths, mux)

	yaml := `
- path: /yt
  url: https://youtube.com
- path: /tf-doc
  url: https://www.tensorflow.org/api_docs
`
	yamlHandler, err := short.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", yamlHandler)

}
