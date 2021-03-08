package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo item",
	Run: func(cmd *cobra.Command, args []string) {
		RemoveTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

// RemoveTodo remove todo item from list of todos
func RemoveTodo(args []string) {
	if len(args) != 1 {
		log.Fatal("Error: Remove requires exactly ONE argument.")
	}

	i, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		fmt.Println("Error: Argument needs to be an integer")
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/remove/%d", AppBaseURL, i),
		nil)

	if err != nil {
		fmt.Println("Error: Failed to execute command. Reason: ")
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error: Failed to execute command. Reason: ")
		log.Fatal(err)
	}

	//Check the HTTP response
	resData, err := ioutil.ReadAll(resp.Body)
	var response Response
	json.Unmarshal(resData, &response)
	if response.Status == "ok" {
		fmt.Println("TODO removed successfully!")
	} else {
		fmt.Printf("Could not remove TODO on the server side. Reason: %s\n", response.ErrorContent)
	}
}
