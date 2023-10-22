package config

import (
	"log"
	"os/exec"
)

func OpenUrlInBrowser(url string) {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
