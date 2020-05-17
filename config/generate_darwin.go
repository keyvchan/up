// Package config contains related stuff for configuration. Darwin
package config

import (
	"log"
	"os"
)

func Generate() {

	HOME = os.Getenv("HOME")
	ConfigPath = HOME + "/.config/up/config.json"

	// check configuration dir exist
	_, err := os.Stat(HOME + "/.config/up")
	if os.IsNotExist(err) {
		err := os.Mkdir(HOME+"/.config/up", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	// check configuration exist
	_, err = os.Stat(ConfigPath)
	if os.IsNotExist(err) {

		f, err := os.OpenFile(ConfigPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.WriteString(Template)
		if err != nil {
			log.Fatal(err)
		}
	}

}
