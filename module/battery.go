package module

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Battery module provides information for a single battery device
type Battery struct {
	Device string // the battery device to query
}

// Status returns the battery capacity percentage
func (b *Battery) Status() string {
	percent := b.batteryInfo("capacity")
	plug := b.plugInfo()
	return fmt.Sprintf("%s%%%s", percent, plug)
}

func (b *Battery) plugInfo() string {
	status := b.batteryInfo("status")
	if status == "Charging" {
		return " "
	}
	return ""
}

func (b *Battery) batteryInfo(attrib string) string {
	const pathFormat = "/sys/class/power_supply/%s/%s"
	path := fmt.Sprintf(pathFormat, b.Device, attrib)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error while reading battery", err)
		return ""
	}
	return strings.Trim(string(d), " \n\t")
}
