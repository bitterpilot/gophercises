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

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTpl))
}

var tpl *template.Template

var defaultHandlerTpl = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
	<section class="page">
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		<ul>
		{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
		{{end}}
		</ul>
	</section>
</body>
<style>
	body {
		font-family: helvetica, arial;
	}
	h1 {
		text-align:center;
		position:relative;
	}
	.page {
		width: 80%;
		max-width: 500px;
		margin: auto;
		margin-top: 40px;
		margin-bottom: 40px;
		padding: 80px;
		background: #FFFCF6;
		border: 1px solid #eee;
		box-shadow: 0 10px 6px -6px #777;
	}
	ul {
		border-top: 1px dotted #ccc;
		padding: 10px 0 0 0;
		-webkit-padding-start: 0;
	}
	li {
		padding-top: 10px;
	}
	a,
	a:visited {
		text-decoration: none;
		color: #6295b5;
	}
	a:active,
	a:hover {
		color: #7792a2;
	}
	p {
		text-indent: 1em;
	}
	</style>
</html>
`

type HandlerOptions func(h *handler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *handler) {
		h.t = t
	}
}

func NewHandler(s Story, opts ...HandlerOptions) http.Handler {
	h := handler{
		s: s,
		t: tpl,
	}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	// remove leading /
	path = path[1:]

	// check if path is a actually a chapter and execute if true
	if chapter, ok := h.s[path]; ok {
		if err := h.t.Execute(w, chapter); err != nil {
			log.Printf("%v\n", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
