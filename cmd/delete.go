/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/damascenov/ursave/config"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a saved URL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		name, _ := cmd.Flags().GetString("name")

		if name == "" {
		    fmt.Println("Name of the URL to delete is required")
		    return
		}

		err := config.DeleteURL(name)
		if err != nil {
		    fmt.Println("Error deleting URL:", err)
		} else {
		    fmt.Printf("URL '%s' deleted successfully.\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("name", "n", "", "Name of the url")
}
