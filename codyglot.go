package main

import (
	"log"

	"github.com/Zenika/codyglot/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
