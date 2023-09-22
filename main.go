package main

import (
	"okki-status/config"
	"okki-status/core"
	"os"
)

var configLocations = []string{
	"$XDG_CONFIG_HOME/okki-status/config.yaml",
	"$HOME/.config/okki-status/config.yaml",
	"/etc/okki-status/config.yaml",
}

func main() {
	cfg := mustLoadConfig()
	bar := mustParse(cfg)
	panic(bar.Start())
}

func mustParse(cfg *os.File) *core.Bar {
	bar, err := config.Parse(cfg)
	if err != nil {
		panic(err)
	}
	return bar
}

func mustLoadConfig() *os.File {
	for _, location := range configLocations {
		filepath := os.ExpandEnv(location)
		f, err := os.Open(filepath)
		if err == nil {
			return f
		}
	}
	panic("config file not found")
}
