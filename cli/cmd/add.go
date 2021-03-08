package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		AddTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}


// AddTodo command
func AddTodo(args []string) {

	if len(args) != 1 {
		log.Fatal("Error: Add requires exactly ONE argument.")
	}

	do := Do{Content: args[0]}

	todoJSON, err := json.Marshal(do)

	if err != nil {
		fmt.Println("Error adding TODO. Reason: ")
		log.Fatal(err)
	}

	_, err = http.Post(
		fmt.Sprintf("%s/add", AppBaseURL),
		"application/json",
		bytes.NewBuffer(todoJSON))

	if err != nil {
		fmt.Println("Error adding TODO. Reason: ")
		log.Fatal(err)
	}

	fmt.Println("TODO added successfully!")
}
