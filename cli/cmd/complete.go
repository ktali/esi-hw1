package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo as complete",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")
		completeTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func completeTodo(args []string) {

	if len(args) > 1 {
		fmt.Println("Error: Too many args.")
	}

	i, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	do := Do{Id: i, IsCompleted: true}

	todoJSON, err := json.Marshal(do)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/complete", AppBaseURL),
		bytes.NewBuffer(todoJSON))

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	_, err = client.Do(req)

	if err != nil {
		// handle error
		log.Fatal(err)
	}
}
