package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var AppBaseURL = "http://localhost:8081"

type Response struct {
	Status       string `json:"status"`
	ErrorContent string `json:"error"`
}

type Do struct {
	Id          int64  `json:"id"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed"`
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A simple cli todo tool",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	if os.Getenv("TODO_SERVER_HOST") != "" {
		AppBaseURL = os.Getenv("TODO_SERVER_HOST")
	}
}
