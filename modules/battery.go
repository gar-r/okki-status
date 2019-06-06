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

	// Display strings for status
	Charging, Discharging string

	// Display strings for capacity levels
	Full, High, Normal, Low, Critical, Unknown string
}

// Status returns the battery status string
func (b *Battery) Status() string {
	return strings.Trim(fmt.Sprintf("%s %s %s",
		b.status(),
		b.capacityLevel(),
		b.capacity()), " ")
}

func (b *Battery) capacity() string {
	capacity := queryFn(b.Battery, "capacity")
	return fmt.Sprintf("%s%%", capacity)
}

func (b *Battery) capacityLevel() string {
	level := queryFn(b.Battery, "capacity_level")
	switch level {
	case "Full":
		return b.Full
	case "High":
		return b.High
	case "Normal":
		return b.Normal
	case "Low":
		return b.Low
	case "Critical":
		return b.Critical
	default:
		return b.Unknown
	}
}

func (b *Battery) status() string {
	status := queryFn(b.Battery, "status")
	switch status {
	case "Charging":
		return b.Charging
	case "Discharging":
		return b.Discharging
	default:
		return ""
	}
}

var queryFn = func(battery, attrib string) string {
	const pathFormat = "/sys/class/power_supply/%s/%s"
	path := fmt.Sprintf(pathFormat, battery, attrib)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error while reading battery capacity", err)
		return "error"
	}
	return string(d)
}
