package sinks

import (
	"log"
	"os/exec"
)

//Xroot is a sink that sets the status on the X root window
type Xroot struct {
}

// Accept acceps the status information
func (*Xroot) Accept(status string) {
	err := exec.Command("xsetroot", "-name", status).Run()
	if err != nil {
		log.Println(err)
	}
}
