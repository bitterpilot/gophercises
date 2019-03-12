package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bitterpilot/gophercises/cyoa"
)

func main() {
	filename := flag.String("file", "./gopher.json", "The JSON file with the cyoa story")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *filename)

	// Open file
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	//decode file
	Story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", Story)
}
