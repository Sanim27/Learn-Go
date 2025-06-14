/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks to do.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
		fmt.Println("These are the todos for you !!!")
		fmt.Println()
		initDB()
		get_list()
	},
}

func get_list() {
	stmt, err := db.Prepare("SELECT * from task")
	if err != nil {
		log.Fatal("Prepare failed: ", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal("Query failed: ", err)
	}
	defer rows.Close()

	fmt.Printf("%-5s | %-50s | %-10s\n", "ID", "Task", "Status")
	fmt.Println("--------------------------------------------------------------------")

	for rows.Next() {
		var id int
		var work, status string

		err := rows.Scan(&id, &work, &status)
		if err != nil {
			log.Fatal("Row scan failed: ", err)
		}
		fmt.Printf("%-5d | %-50s | %-10s\n", id, work, status)
	}

	//check if any errors occurred during iteration
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func init() {
	rootCmd.AddCommand(listCmd)
}
