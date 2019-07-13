package modules

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Battery provides battery related status
type Battery struct {
	Margin

	// Name of the battery to query
	Battery string

	// Maps different statuses to custom strings
	StatusMap map[string]string
}

// Status returns the battery status string
func (b *Battery) Status() string {
	return b.Format(fmt.Sprintf("%s %s", b.status(), b.capacity()))
}

func (b *Battery) capacity() string {
	capacity := b.getInfo("capacity")
	return fmt.Sprintf("%s%%", capacity)
}

func (b *Battery) status() string {
	status := b.getInfo("status")
	return b.StatusMap[status]
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
