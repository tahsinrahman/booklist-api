package main

import (
	"log"

	"github.com/tahsinrahman/booklist-api/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
