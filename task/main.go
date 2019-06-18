package main

import (
	"log"
	"path/filepath"

	"github.com/bitterpilot/gophercises/task/cmd"
	"github.com/bitterpilot/gophercises/task/db"
)

func main() {
	// get file path for bolt db and connect
	home := "."
	// home, err := homedir.Dir()   //"github.com/mitchellh/go-homedir"
	// if err != nil {
	// 	log.Fatalf("User home directory not found")
	// }
	dbPath := filepath.Join(home, "tasks.boltDB")
	err := db.Connect(dbPath)
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	cmd.RootCmd.Execute()
}
