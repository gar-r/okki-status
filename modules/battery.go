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

	// Display glyph for charging status
	Charging string

	// Display glyph for discharging status
	Discharging string
}

// Status returns the battery status string
func (b *Battery) Status() string {
	return fmt.Sprintf("%s%s", b.status(), b.capacity())
}

func (b *Battery) capacity() string {
	capacity := b.getInfo("capacity")
	return fmt.Sprintf("%s%%", capacity)
}

func (b *Battery) status() string {
	status := b.getInfo("status")
	if status == "Charging" {
		return b.Charging
	}
	return b.Discharging
}

func (b *Battery) getInfo(attrib string) string {
	const pathFormat = "/sys/class/power_supply/%s/%s"
	path := fmt.Sprintf(pathFormat, b.Battery, attrib)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error while reading battery capacity", err)
		return ""
	}
	return strings.Trim(string(d), " \n\t")
}
