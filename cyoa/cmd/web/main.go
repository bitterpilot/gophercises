package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bitterpilot/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "The port to start CYOA web application on")
	filename := flag.String("file", "./gopher.json", "The JSON file with the cyoa story")
	flag.Parse()
	fmt.Printf("Using the file %s\n", *filename)

	// Open file
	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	//decode file
	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	// start webserver
	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
