/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"

	"github.com/damascenov/ursave/config"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a url",
	Long:  ``,
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

func PromptAdd() {
	namePrompt := promptui.Prompt{
		Label: "Name",
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

	url, err := urlPrompt.Run()
	if err != nil {
		fmt.Printf("Failed %v\n", err)
		return
	}

	err = config.AddUrl(name, url)
	if err != nil {
		fmt.Println("Error adding URL:", err)
	} else {
		fmt.Printf("URL '%s' added successfully.\n", name)
	}

}
