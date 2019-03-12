package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var Story Story
	if err := d.Decode(&Story); err != nil {
		return nil, err
	}
	return Story, nil
}

var DefultHandlerTpl = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
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
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(DefultHandlerTpl))

	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	// remove leading /
	path = path[1:]

	// check if path is a actually a chapter and execute if true
	if chapter, ok := h.s[path]; ok {
		if err := tpl.Execute(w, chapter); err != nil {
			log.Printf("%v\n", err)
		    http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
