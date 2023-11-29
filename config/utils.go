package config

import (
	"log"
	"net/url"
	"os/exec"
)

func OpenUrlInBrowser(url string) {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func IsValidUrl(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme!= "" && u.Host!= ""
}
