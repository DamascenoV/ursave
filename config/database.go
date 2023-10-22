package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Url struct {
	ID   int
	Name string
	Url  string
}

func InitializeDB() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(homeDir, "ursave.db")

	var errOpen error
	db, errOpen = sql.Open("sqlite3", dbPath)
	if errOpen != nil {
		log.Fatal(errOpen)
	}

	createTable := `
	    CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY,
		name TEXT,
		url TEXT
	    );`

	_, errCreate := db.Exec(createTable)
	if errCreate != nil {
		log.Fatal(errCreate)
	}
}

func GetDB() *sql.DB {
	return db
}

func GetUrls() []Url {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
	}

	query := "SELECT id, name, url FROM urls"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var urls []Url

	for rows.Next() {
		var url Url
		err := rows.Scan(&url.ID, &url.Name, &url.Url)
		if err != nil {
			log.Fatal(err)
		}
		urls = append(urls, url)
	}

	return urls
}

func GetUrl(name string) (Url, error) {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
	}

	query := "SELECT id, name, url FROM urls WHERE name = ?"
	row := db.QueryRow(query, name)
	var url Url
	err := row.Scan(&url.ID, &url.Name, &url.Url)
	if err != nil {
		return url, err
	}

	return url, nil
}

func AddUrl(name string, url string) error {
	if db == nil {
		log.Fatal("Database connection is not initialized.")
	}

	query := "INSERT INTO urls (name, url) VALUES (?, ?)"
	_, err := db.Exec(query, name, url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("URL added:", name, url)

	return nil
}

func DeleteURL(name string) error {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}

	deleteQuery := "DELETE FROM urls WHERE name = ?"
	_, err := db.Exec(deleteQuery, name)
	if err != nil {
		return err
	}

	return nil
}

func EditURL(name, newURL string) error {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}

	updateQuery := "UPDATE urls SET url = ? WHERE name = ?"
	_, err := db.Exec(updateQuery, newURL, name)
	if err != nil {
		return err
	}

	return nil
}
