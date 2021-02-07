package main

import (
	"log"
	"os"
)

func main() {
	modules := os.Args[1:]
	for _, module := range modules {
		refresh(module)
	}
}

func refresh(module string) {
	log.Printf("%s refreshing", module)
	err := sendRefreshRequest(module)
	if err != nil {
		log.Printf("%s not refreshed: %v", module, err)
	}
}
