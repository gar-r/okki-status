package providers

import (
	"bitbucket.org/dargzero/okki-status/core"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Battery struct {
	// Name of the battery to query
	Battery string
}

func (b *Battery) GetStatus() string {
	percent := batteryInfo(b.Battery, "capacity")
	return fmt.Sprintf("%s%%", percent)
}

type BatteryIconProvider struct {
	core.ThresholdIcon

	// Name of the battery to query
	Battery string

	Charging string
}

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
