package main

import (
	"flag"
	"fmt"
	"os"
	"rem-lag/cyoa/teller"
)

func main() {
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

	fmt.Printf("%+v\n", story)
}
