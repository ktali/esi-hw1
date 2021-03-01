package cmd

import (
	"github.com/spf13/cobra"
)

const (
	AppBaseURL = "http://localhost:8081"
)

type Do struct {
	Id          int64  `json:"id"`
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed"`
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo-cli.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
