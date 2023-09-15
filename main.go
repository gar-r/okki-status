package main

import (
	"okki-status/config"
	"os"
)

func main() {
	f, err := os.Open("etc/example.yaml")
	if err != nil {
		panic(err)
	}
	bar, err := config.Parse(f)
	if err != nil {
		panic(err)
	}
	bar.Start()
}
