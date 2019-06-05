package main

import (
	"fmt"
	"strings"
)

func main() {
	config := NewConfig()
	status := strings.Builder{}
	for _, entry := range config.Entries {
		status.WriteString(getFormattedData(entry))
	}
	fmt.Println(status.String())
}

func getFormattedData(entry Entry) string {
	data := entry.provider.GetData()
	if entry.format != "" {
		data = fmt.Sprintf(entry.format, data)
	}
	return data
}
