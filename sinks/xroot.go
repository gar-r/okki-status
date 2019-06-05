package sinks

import (
	"log"
	"os/exec"
)

type Xroot struct {
}

func (*Xroot) Accept(status string) {
	err := exec.Command("xsetroot", "-name", status).Run()
	if err != nil {
		log.Println(err)
	}
}
