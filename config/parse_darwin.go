package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func Parse() []Item {

	fileBytes, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	var Items []Item

	err = json.Unmarshal(fileBytes, &Items)
	if err != nil {
		log.Fatal(err)
	}

	return Items

}
