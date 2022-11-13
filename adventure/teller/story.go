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

// Implements http.Handler interface
// so it can be used with http package
type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

// Request endpoint should be chapter title
// Return requested story chapter with requested template
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathFn(r)
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

// Function type user can pass to NewHandler to
// set struct fields
type HandlerOptions func(h *handler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFn(fn func(r *http.Request) string) HandlerOptions {
	return func(h *handler) {
		h.pathFn = fn
	}
}

func defaultPath(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	return path[1:]
}

// Return handler struct with and desired Handler option funcs
// that will be applied to the handler struct
// Returned stuct used for ListenAndServe
func NewHandler(s Story, opts ...HandlerOptions) http.Handler {
	h := handler{s, tmpl, defaultPath}

	for _, opt := range opts {
		opt(&h)
	}
	return h
}

var tmpl *template.Template

// Pass sting of html template to parse into
// type that can be used by the handler
func ParseTemplate(temp string) *template.Template {
	t := template.Must(template.New("").Parse(temp))
	return t
}

func init() {
	tmpl = ParseTemplate(defaultHandlerTemp)
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
