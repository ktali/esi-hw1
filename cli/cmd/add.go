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
		addTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTodo(args []string) {

	if len(args) > 1 {
		fmt.Println("Error: Too many args.")
	}

	do := Do{Content: args[0]}

	todoJSON, err := json.Marshal(do)

	if err != nil {
		log.Fatal(err)
	}

	_, err = http.Post(
		fmt.Sprintf("%s/add", AppBaseURL),
		"application/json",
		bytes.NewBuffer(todoJSON))

	if err != nil {
		log.Fatal(err)
	}
}
