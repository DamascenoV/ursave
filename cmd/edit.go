/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/damascenov/ursave/config"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a saved URL",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		newURL, _ := cmd.Flags().GetString("new-url")

		if name == "" || newURL == "" {
			fmt.Println("Both name and new URL are required")
			return
		}

		err := config.EditURL(name, newURL)
		if err != nil {
			log.Fatal("Error editing URL:", err)
		} else {
			fmt.Printf("URL '%s' edited successfully.\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("name", "n", "", "Name of the URL to edit")
	editCmd.Flags().StringP("new-url", "u", "", "New URL address")
}
