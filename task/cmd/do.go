package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bitterpilot/gophercises/task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse argument:", arg)
				fmt.Println("do requires whole numbers")
			}
			ids = append(ids, id)
		}

		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatalf("doCmd: something went wrong: %v", err)
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Printf("Invalid task number: %d\n", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark %d as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Mark %d as completed.\n", id)

			}
		}

	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
