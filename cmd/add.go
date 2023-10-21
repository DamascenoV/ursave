/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"

	"github.com/damascenov/ursave/config"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a url",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
        url, _ := cmd.Flags().GetString("url")

        if name == "" || url == "" {
            fmt.Println("Both name and URL are required")
            return
        }

        err := config.AddUrl(name, url)
        if err != nil {
            fmt.Println("Error adding URL:", err)
        } else {
            fmt.Printf("URL '%s' added successfully.\n", name)
        }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Name of the url")
	addCmd.Flags().StringP("url", "u", "", "Url of the Address")
}
