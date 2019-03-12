package cyoa

import (
	"encoding/json"
	"io"
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

var DefultHandlerTmpl = `
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