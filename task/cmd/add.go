package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"github.com/bitterpilot/gophercises/task/db"
)

var addCMD = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			log.Fatalf("addCMD: something went wrong: %v", err)
		}
		fmt.Printf("Added \"%s\" to your tasklist\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCMD)
}
