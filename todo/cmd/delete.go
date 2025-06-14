/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Removing a work from the todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please enter the id of the task you want to remove.")
		}
		initDB()
		remove_task_id := args[0]
		remove_task(remove_task_id)

	},
}

func remove_task(task_id string) {
	stmt, err := db.Prepare("DELETE FROM task WHERE ID=?")
	if err != nil {
		log.Fatal("Error in preparation", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(task_id)
	if err != nil {
		log.Fatal("Execution failed: ", err)
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Printf("Task with id %s deleted (%d rows affected). \n", task_id, rowsAffected)
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
