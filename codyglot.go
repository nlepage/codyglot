package main

import (
	"log"

	"github.com/nlepage/codyglot/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
