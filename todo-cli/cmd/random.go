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

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Display a random ToDo",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomCmd()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

//Prints the ToDo item
func getRandomCmd() {
	res := GetRandom()

	var do Do
	json.Unmarshal(res, &do)

	fmt.Println("A random TODO:")
	fmt.Printf("Id: %d\tContent: %s\tDone: %t\n", do.Id, do.Content, do.IsCompleted)
}

//GetRandom fetches a random ToDo item
func GetRandom() []byte {
	resp, err := http.Get(fmt.Sprintf("%s/random", AppBaseURL))

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