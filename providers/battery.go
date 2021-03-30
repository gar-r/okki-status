package providers

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"hu.okki.okki-status/core"
)

// Battery provides status information for a single battery device
type Battery struct {
	Battery string // name of the battery to query
}

// GetStatus returns a string with the battery capacity percentage
func (b *Battery) GetStatus() string {
	percent := batteryInfo(b.Battery, "capacity")
	return fmt.Sprintf("%s%%", percent)
}

// BatteryIconProvider provides a state-aware battery icon
type BatteryIconProvider struct {
	core.ThresholdIcon

	Battery  string // name of the battery to query
	Charging string
}

// GetIcon returns different icons depending of charging/discharging state and percentage
func (b *BatteryIconProvider) GetIcon(status string) string {
	state := batteryInfo(b.Battery, "status")
	if state == "Charging" {
		return b.Charging
	}
	return b.ThresholdIcon.GetIcon(status)
}

func batteryInfo(battery, attrib string) string {
	const pathFormat = "/sys/class/power_supply/%s/%s"
	path := fmt.Sprintf(pathFormat, battery, attrib)
	d, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("error while reading battery capacity", err)
		return ""
	}
	return strings.Trim(string(d), " \n\t")
}
