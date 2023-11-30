package config

import (
	"log"
	"fmt"
	"net/url"
	"os/exec"

	"github.com/koki-develop/go-fzf"
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

func GetSelectedUrl() (Url, error) {
	items := GetUrls()

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
		error := fmt.Errorf("No urls found")
		return Url{}, error
	}

	return items[idxs[0]], nil
}
