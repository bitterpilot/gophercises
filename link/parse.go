package link

import (
	"io"
)

// Link is a simplified version of <a href="..."> link
type Link struct {
	Href string
	Text string
}

// Parse takes a HTML document and returns the links in a simplified struct.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
