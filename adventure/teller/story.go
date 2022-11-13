package teller

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

// Story type contains a chapter title key and parsed json as Chapter struct value
type Story map[string]Chapter

// Struct for json story
// Paragraph is story text, options field stored as slice of options struct
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

// Options field of json contains the prompt for the option (Text)
// and chapter title (Chapter)
type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// Parses json into go Story type
func JsonStory(r io.Reader) (Story, error) {
	var story Story

	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil

}

// Return handler struct with a template and the parsed story
// Returned stuct used for ListenAndServe
func NewHandler(s Story, t *template.Template) http.Handler {
	if t == nil {
		t = tmpl
	}
	return handler{s, t}
}

// Implements http.Handler interface
type handler struct {
	s Story
	t *template.Template
}

// Request endpoint should be chapter title
// Return requested story chapter with requested template
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if ch, ok := h.s[path]; ok {
		err := h.t.Execute(w, ch)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found...", http.StatusNotFound)
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Parse(defaultHandlerTemp))
}

var defaultHandlerTemp = `
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
            <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
    </body>
</html>`
