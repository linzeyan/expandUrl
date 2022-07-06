package main

import (
	"fmt"
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
	var resp, err = expandUrl.Expand(url)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(resp)
}

func main() {
	rootCmd.Flags().StringVarP(&url, "URL", "u", "", "Specify URL to expand")
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
