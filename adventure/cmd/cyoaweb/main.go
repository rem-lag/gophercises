package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"rem-lag/cyoa/teller"
	"strings"
)

// Function to demonstrate usage option functions
func CustomPath(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

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

	temp := teller.ParseTemplate(storyHandlerTemp)

	h := teller.NewHandler(story,
		teller.WithTemplate(temp),
		teller.WithPathFn(CustomPath),
	)
	mux := http.NewServeMux()
	mux.Handle("/story/", h)
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

var storyHandlerTemp = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Choose Your Own Adventure!!</title>
    </head>
    <body>
        <h1>{{.Title}}</h1>
        {{range .Paragraphs}}
            <p>{{.}}</p>
        {{end}}
        <ul>
        {{range .Options}}
            <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`
