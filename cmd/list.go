/*
Copyright © 2023 DAMASCENOV
*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/koki-develop/go-fzf"
	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
		initializeDb()
		getUrls()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func initializeDb() {
	dbPath := "ursave.db"
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func getUrls() {
	query := "SELECT name FROM urls"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var items []string

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, name)
	}

	fzf, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := fzf.Find(items, func(i int) string {
		return items[i]
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println("ESSE É O INDEX", i)
	}
}
