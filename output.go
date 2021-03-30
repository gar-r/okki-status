package main

import (
	"fmt"
	"log"
	"os/exec"
)

var xroot func(string) = func(status string) {
	err := exec.Command("xsetroot", "-name", status).Run()
	if err != nil {
		log.Println(err)
	}
}

var stdout func(string) = func(status string) {
	fmt.Println(status)
}
