/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/damascenov/ursave/config"
	"github.com/manifoldco/promptui"
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

		err := config.EditURL(name, "", newURL)
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

func PromptEdit() {
	selectedURL, err := config.GetSelectedUrl()

	if err != nil {
		log.Fatal("Error getting selected URL:", err)
	}

	fmt.Printf("Editing URL '%s':\n", selectedURL.Name)
	fmt.Printf("Url: %s\n", selectedURL.Url)


	namePrompt := promptui.Prompt{
		Label: "Name",
		Default: selectedURL.Name,
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("Cannot be empty")
			}
			return nil
		},
	}

	name, err := namePrompt.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	urlPrompt := promptui.Prompt{
		Label: "URL",
		Default: selectedURL.Url,
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("Cannot be empty")
			}

			isValid := config.IsValidUrl(input)

			if !isValid {
				return fmt.Errorf("Invalid URL")
			}
			return nil
		},
	}

	newURL, err := urlPrompt.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	err = config.EditURL(selectedURL.Name, name, newURL)
	if err != nil {
		log.Fatal("Error editing URL:", err)
	} else {
		fmt.Printf("URL '%s' edited successfully.\n", name)
	}
}
