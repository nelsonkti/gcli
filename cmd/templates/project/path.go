// Package project
// @Author fzy
// @Date 2022-04-14 15:04:03
package project

import (
	"log"
	"os"
	"path"
)

func HomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	home := path.Join(dir, ".nelsonkti")
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}

	return home
}
