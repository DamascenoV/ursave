/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/damascenov/ursave/config"
)



var rootCmd = &cobra.Command{
	Use:   "UrSave",
	Short: "Url Saver CLI",
	Long: `UrSave is a CLI url saver.
	You can save your favorite urls and open in the browser`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating Database")
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
}


