package link

import (
	"io"
	"strings"

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
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		link := Link{
			Href: strings.TrimSpace(node.Attr[0].Val),
			Text: strings.TrimSpace(node.FirstChild.Data),
		}
		links = append(links, link)
	}
	return links, nil
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
