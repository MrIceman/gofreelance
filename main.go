package main

import (
	"gofreelance/data"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	handler := NewHandler(data.NewDefaultEntriesHandler())

	if len(args) == 0 {
		log.Println("no arguments passed in, use 'help' to see available commands")
		return
	}
	if err := handler.Handle(args...); err != nil {
		log.Fatalf("error: %s", err)
	}
}
