package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Battery provides battery related status
type Battery struct {
	// Name of the battery to query
	Battery string

	// Display string for charging status
	Charging string
}

// Status returns the battery status string
func (b *Battery) Status() string {
	return strings.Trim(fmt.Sprintf("%s%s",
		b.status(),
		b.capacity()), " ")
}

func (b *Battery) capacity() string {
	capacity := batteryFn(b.Battery, "capacity")
	return fmt.Sprintf("%s%%", capacity)
}

func (b *Battery) status() string {
	status := batteryFn(b.Battery, "status")
	if status == "Charging" {
		return b.Charging
	}
	return ""
}

var batteryFn = func(battery, attrib string) string {
	const pathFormat = "/sys/class/power_supply/%s/%s"
	path := fmt.Sprintf(pathFormat, battery, attrib)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error while reading battery capacity", err)
		return ""
	}
	return strings.Trim(string(d), " \n\t")
}
