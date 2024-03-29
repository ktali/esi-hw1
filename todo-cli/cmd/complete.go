package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a todo as complete",
	Run: func(cmd *cobra.Command, args []string) {
		CompleteTodo(args)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func CompleteTodo(args []string) {

	if len(args) != 1 {
		log.Fatal("Error: Complete requires exactly ONE argument.")
	}

	i, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		fmt.Println("Error: Argument needs to be an integer")
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("%s/complete/%d", AppBaseURL, i),
		bytes.NewBuffer([]byte{}))

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
		fmt.Println("TODO completed successfully!")
	} else {
		fmt.Printf("Could not complete TODO on the server side. Reason: %s\n", response.ErrorContent)
	}
}
