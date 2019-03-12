package main

import (
	"github.com/bitterpilot/gophercises/cyoa"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
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
	d := json.NewDecoder(f)
	var Story cyoa.Story
	if err := d.Decode(&Story); err != nil{
		log.Fatalf("%s\n", err)
	} 
	fmt.Printf("%+v", Story)
}
