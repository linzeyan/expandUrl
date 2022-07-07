package main

import (
	"log"
	"os"

	"github.com/linzeyan/expandUrl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expandUrl",
	Short: "Expand shorten URL.",
	Long:  "Expand shorten URL.",
	Args:  cobra.OnlyValidArgs,
	Run:   rootCmdRun,
}
var url string

func rootCmdRun(_ *cobra.Command, _ []string) {
	_, err := expandUrl.Expand(url)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Specify URL to expand")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
