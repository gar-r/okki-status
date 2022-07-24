package config

import (
	"time"

	"okki-status/core"
	"okki-status/module"
)

// block declarations

var clock = &core.Block{
	Name:   "clock",
	Prefix: " ",
	Module: &module.Clock{Format: "2006-01-02 15:04:05"},
}

var layout = &core.Block{
	Name:   "layout",
	Prefix: " ",
	Module: &module.KeyboardSw{
		Identifier: "1:1:AT_Translated_Set_2_keyboard",
	},
}

var battery = &core.Block{
	Name:   "battery",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.Battery{
			Device: "BAT0",
		},
		time.Minute,
	),
}

var brightness = &core.Block{
	Name:   "brightness",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.Brightness{},
		5*time.Second,
	),
}

var volume = &core.Block{
	Name:   "volume",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.Volume{},
		5*time.Second,
	),
}

var ram = &core.Block{
	Name:   "ram",
	Prefix: " ",
	Module: &module.RAM{},
}

var wifi = &core.Block{
	Name:   "wifi",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.WiFi{
			Device: "wlan0",
		},
		5*time.Second,
	),
}

var updates = &core.Block{
	Name:   "updates",
	Prefix: " ",
	Module: core.NewCachingModule(
		&module.Updates{
			Command:         "/usr/bin/checkupdates",
			IgnoreExitError: true,
		},
		time.Hour,
	),
}
