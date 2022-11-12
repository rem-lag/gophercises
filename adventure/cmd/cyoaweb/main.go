package main

import (
	"encoding/json"
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

	var story teller.Story

	d := json.NewDecoder(f)
	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
