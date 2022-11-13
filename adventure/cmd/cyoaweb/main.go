package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"rem-lag/cyoa/teller"
)

func main() {
	port := flag.Int("port", 8086, "port to start cyoa server on")
	file := flag.String("file", "story.json", "JSON containing story and options")
	flag.Parse()
	fmt.Printf("Using story %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := teller.JsonStory(f)
	if err != nil {
		panic(err)
	}

	h := teller.NewHandler(story)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
