/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"log"
	"os"

	"github.com/damascenov/ursave/config"
	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "UrSave",
	Short: "Url Saver CLI",
	Long: `UrSave is a CLI url saver.
	You can save your favorite urls and open in the browser`,
	Run: func(cmd *cobra.Command, args []string) {
		open, _ := cmd.Flags().GetString("open")

		if open == "" {
			option, err := getOptions(cmd)
			if err != nil {
				log.Fatal(err)
			}

			runOption(option)
		}

		url, err := config.GetUrl(open)
		if err != nil {
			log.Fatal("Record not found in UrSave")
			return
		}

		config.OpenUrlInBrowser(url.Url)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	config.InitializeDB()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ursave.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("open", "o", "", "Open the url in the browser")
}


func getOptions(cmd *cobra.Command) (string, error) {
	items := []string{"add", "edit", "list", "delete"}

	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(items, func(i int) string { return items[i] })
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		return items[i], nil
	}

	return "", err
}

func runOption(option string) {
	if option == "add" {
		PromptAdd()
		return
	}
	
	if option == "list" {
		GetUrls()
		return
	}

	if option == "delete" {
		PromptDelete()
		return
	}

	println("Invalid")
	return
}
