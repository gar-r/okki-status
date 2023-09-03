package main

import (
	"fmt"
	"okki-status/config"
	"os"
)

func main() {
	f, err := os.Open("etc/example.yaml")
	if err != nil {
		panic(err)
	}
	conf, err := config.Read(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(conf.Modules[0].Status())
}
