/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"

	"github.com/damascenov/ursave/config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved urls",
	Run: func(cmd *cobra.Command, args []string) {
		GetUrls()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func GetUrls() {
	selectedUrl, err := config.GetSelectedUrl()

	if err != nil {
		fmt.Println("Error getting selected URL:", err)
		return
	}

	config.OpenUrlInBrowser(selectedUrl.Url)
}
