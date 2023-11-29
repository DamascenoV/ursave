/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/damascenov/ursave/config"
	"github.com/koki-develop/go-fzf"
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
	items := config.GetUrls()

	fzf, err := fzf.New(
		fzf.WithInputPosition(fzf.InputPositionBottom),
	)
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
	config.OpenUrlInBrowser(selectedUrl)
}
