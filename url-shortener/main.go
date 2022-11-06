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
	tempPaths := map[string]string{
		"url-1": "https://website.com",
	}

	mapHandler := short.MapHandler(tempPaths, mux)

}
