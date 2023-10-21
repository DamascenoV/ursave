/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/damascenov/ursave/config"
	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getUrls()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func getUrls() {
	items := config.GetUrls()

	fzf, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := fzf.Find(items, func(i int) string {
		return items[i].Name
	})
	if err != nil {
		log.Fatal(err)
	}

	if len(idxs) == 0 {
		fmt.Println("No urls found")
		return
	}

	selectedUrl := items[idxs[0]].Url
	openUrlInBrowser(selectedUrl)
}

func openUrlInBrowser(url string) {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	
}
