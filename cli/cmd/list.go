package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display a list of todos",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		listDos()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listDos() {
	resp, err := http.Get(AppBaseURL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resData))
}