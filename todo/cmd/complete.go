/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Use this to mark a task done.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Enter an id for a task that you've completed.")
		}
		id := args[0]
		initDB()
		markCompleted(id)
	},
}

func markCompleted(id string) {
	stmt, err := db.Prepare("Update task set status = ? where id=?")
	if err != nil {
		log.Fatal("DB cannot be prepared", err)
	}
	res, err := stmt.Exec("Completed", id)
	if err != nil {
		log.Fatal("Execution failed:", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		fmt.Printf("No task found with ID %s.\n", id)
	} else {
		fmt.Printf("Task %s marked as completed.\n", id)
	}
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
