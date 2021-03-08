package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of todos",
	Run: func(cmd *cobra.Command, args []string) {
		listDosCmd()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

//Prints the todo list
func listDosCmd() {
	res := ListDos()

	var dos []Do
	json.Unmarshal(res, &dos)

	fmt.Println("Your list of TODOs:")
	for _, do := range dos {
		fmt.Printf("Id: %d\tContent: %s\tDone: %t\n", do.Id, do.Content, do.IsCompleted)
	}
}

// ListDos fetches all todo items
func ListDos() []byte {
	resp, err := http.Get(AppBaseURL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resData
}
