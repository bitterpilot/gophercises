package urlshort

var buildMapCases = []struct {
	input    []pathURL
	expected map[string]string
}{
	{
		[]pathURL{
			{"/urlshort", "https://github.com/gophercises/urlshort"},
			{"/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution"},
		},
		map[string]string{"/urlshort": "https://github.com/gophercises/urlshort",
			"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution"},
	},
}
