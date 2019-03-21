package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/bitterpilot/gophercises/link"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	flagURL := flag.String("URL", "https://gophercises.com", "The URL to start the mapping from")
	flagMaxDepth := flag.Int("depth", 12, "how meany links 'down' to look")
	flag.Parse()

	pages := bfs(*flagURL, *flagMaxDepth)

	var toXML urlset
	toXML.Xmlns = xmlns
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	fmt.Print(xml.Header)
	if err := enc.Encode(toXML); err != nil {
		fmt.Println(err)
	}
}

func bfs(urlStr string, maxD int) []string {
	// keep track of what has be visited
	// key will be url and value will be empty struct(empty struct has the
	// lowest memory usage of all the types)
	// https://dave.cheney.net/2014/03/25/the-empty-struct
	seen := make(map[string]struct{})
	var currentQ map[string]struct{}
	nextQ := map[string]struct{}{
		urlStr: struct{}{},
	}

	for i := 0; i <= maxD; i++ {
		currentQ = nextQ
		nextQ = make(map[string]struct{})

		for url := range currentQ {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}
			for _, link := range getPage(url) {
				nextQ[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url := range seen {
		ret = append(ret, url)
	}
	return ret
}

func getPage(u string) []string {
	resp, err := http.Get(u)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// set base URL
	// Get url from response incase of redirects
	redURL := resp.Request.URL
	baseURL := url.URL{
		Scheme: redURL.Scheme,
		Host:   redURL.Host,
	}
	base := baseURL.String()
	links, err := hrefs(resp.Body, base)
	if err != nil {
		log.Println(err)
	}

	return filter(base, links)
}

func hrefs(r io.Reader, base string) ([]string, error) {
	// Find links
	li, err := link.Parse(r)
	if err != nil {
		return nil, err
	}

	// remove invalid links (#, mailto)
	var ret []string

	for _, l := range li {
		switch {
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		}
	}
	return ret, nil
}

func filter(base string, links []string) []string {
	var ret []string
	for _, l := range links {
		if strings.HasPrefix(l, base) {
			ret = append(ret, l)
		}
	}
	return ret
}

func validateURL(u *url.URL) {
	if u.String() == "" {
		fmt.Fprintf(os.Stderr, "URL Required\n\tUse -URL to pass in a valid URL\n")
		os.Exit(1)
	}
	// TODO: check URL is valid
	switch {
	case u.Scheme == "":
		fmt.Fprintf(os.Stderr,
			"Valid URL Required\n\tMissing Scheme( http:// or https:// ) \n")
		os.Exit(1)
	case u.Scheme == "http", u.Scheme == "https":
		return
	default:
		fmt.Fprintf(os.Stderr,
			"Valid URL Required\n\tUse -URL to pass in a valid URL\nIn the form of")
		os.Exit(1)
	}
}
