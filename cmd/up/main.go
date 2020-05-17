package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/keyvchan/up/config"
)

func main() {

	config.Generate()
	Items := config.Parse()

	for _, item := range Items {
		cmd := exec.Command(item.Command, item.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}
