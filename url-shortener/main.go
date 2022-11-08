package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
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
	yamlFile := flag.String("yaml", "urls.yaml", "a yaml file containing short path and URL")
	flag.Parse()

	file, err := os.ReadFile(*yamlFile)
	if err != nil {
		fmt.Printf("Faile to open %s\n", *yamlFile)
		os.Exit(1)
	}

	mux := defaultMux()
	mapPaths := map[string]string{
		"/rem-lag":  "https://github.com/rem-lag",
		"/cv-stack": "https://stats.stackexchange.com",
	}

	mapHandler := short.MapHandler(mapPaths, mux)

	// 	yaml := `
	// - path: /yt
	//   url: https://youtube.com
	// - path: /tf-doc
	//   url: https://www.tensorflow.org/api_docs
	// `
	yamlHandler, err := short.YAMLHandler([]byte(file), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", yamlHandler)

}
