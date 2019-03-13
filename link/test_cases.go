package link

import (
	"fmt"
	"io"
	"os"
)

type args struct {
	r io.Reader
}

var tests = []struct {
	name    string
	args    args
	want    []Link
	wantErr bool
}{
	{
		name: "ex1",
		args: args{
			r: open("./ex1.html"),
		},
		want: []Link{
			{
				Href: "/other-page",
				Text: "A link to another page",
			},
		},
		wantErr: false,
	}, {
		name: "ex2",
		args: args{
			r: open("./ex2.html"),
		},
		want: []Link{
			{
				Href: "https://www.twitter.com/joncalhoun",
				Text: "Check me out on twitter",
			},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			},
		},
		wantErr: false,
	}, {
		name: "ex3",
		args: args{
			r: open("./ex3.html"),
		},
		want: []Link{
			{
				Href: "#",
				Text: "Login",
			},
			{
				Href: "/lost",
				Text: "Lost? Need help?",
			},
			{
				Href: "https://twitter.com/marcusolsson",
				Text: "@marcusolsson",
			},
		},
		wantErr: false,
	}, {
		name: "ex4",
		args: args{
			r: open("./ex4.html"),
		},
		want: []Link{
			{
				Href: "/dog-cat",
				Text: "dog cat",
			},
		},
		wantErr: false,
	},
}

func open(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}
