package main

import (
	"okki-status/config"
	"os"
)

func main() {
	f, err := os.Open("/home/garric/projects/okki-status/etc/example.yaml")
	if err != nil {
		panic(err)
	}
	bar, err := config.Parse(f)
	if err != nil {
		panic(err)
	}
	err = bar.Start()
	if err != nil {
		panic(err)
	}
}
