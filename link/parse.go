package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link is a simplified version of <a href="..."> link
type Link struct {
	Href string
	Text string
}

// Parse takes a HTML document and returns the links in a simplified struct.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	fmt.Println(padding, n.Data)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, "  ")
	}
}
