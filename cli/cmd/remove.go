package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo item",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		removeTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeTodo(args []string) {
	if len(args) > 1 {
		fmt.Println("Error: Too many args.")
	}

	i, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/delete%d", AppBaseURL, i),
		nil)

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
