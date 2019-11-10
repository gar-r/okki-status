package output

import (
	"log"
	"os/exec"
)

//XRoot is a sink that sets the status on the X root window
type XRoot struct {
	d debouncer
}

// Accept accepts the status information
func (x *XRoot) Accept(status string) {
	x.d.invoke(func() {
		err := exec.Command("xsetroot", "-name", status).Run()
		if err != nil {
			log.Println(err)
		}
	})
}
