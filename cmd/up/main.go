package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/keyvchan/up/config"
)

func main() {

	config.Generate()
	Items := config.Parse()

	for _, item := range Items {
		fmt.Println("\u001b[32;1m"+"\u279c", " "+item.Kind+":", item.Name, "\u001b[0m")
		cmd := exec.Command(item.Command, item.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}

}
