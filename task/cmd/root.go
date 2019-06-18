package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd makes the root command for the task list.
var RootCmd = &cobra.Command{
	Use:   "Task",
	Short: "Task is a CLI task manager",
}
