/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add tasks",
	Short: "Add a new task.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a task.")
			return
		}
		task := args[0]
		initDB()
		insertTask(task)

		fmt.Println("Task added:", task)
	},
}

func insertTask(work string) {
	stmt, err := db.Prepare("INSERT INTO task(work, status) VALUES(?, ?)")
	if err != nil {
		log.Fatal("Prepare failed:", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(work, "pending")
	if err != nil {
		log.Fatal("Exec Failed:", err)
	}

	id, _ := res.LastInsertId()
	fmt.Printf("Task added with ID %d\n", id)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
