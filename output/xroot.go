package output

import (
	"log"
	"os/exec"
)

//XRoot is a sink that sets the status on the X root window
type XRoot struct {
}

// Accept accepts the status information
func (*XRoot) Accept(status string) {
	err := exec.Command("xsetroot", "-name", status).Run()
	if err != nil {
		log.Println(err)
	}
}
