package main

import (
	"time"

	"hu.okki.okki-status/core"
	"hu.okki.okki-status/providers"
)

var addr = ":12650"

var modules = []core.Module{
	updates,
	wifi,
	ram,
	volume,
	brightness,
	battery,
	layout,
	clock,
}

var updates = core.Module{
	Name: "updates",
	Status: &providers.Updates{
		Command:         "/usr/bin/checkupdates",
		IgnoreExitError: true,
	},
	Icon:       &core.StaticIcon{Icon: " "},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    10 * time.Minute,
}

var wifi = core.Module{
	Name:       "wifi",
	Status:     &providers.WiFi{Device: "wlp0s20f3"},
	Icon:       &core.StaticIcon{Icon: "  "},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    5 * time.Second,
}

var ram = core.Module{
	Name:       "ram",
	Status:     &providers.RAM{},
	Icon:       &core.StaticIcon{Icon: "  "},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    15 * time.Second,
}

var volume = core.Module{
	Name:   "volume",
	Status: &providers.Volume{},
	Icon: &core.ThresholdIcon{
		StatusConverterFn: core.NtoI,
		Thresholds: []core.Threshold{
			{Value: 50, Icon: " "},
			{Value: 25, Icon: " "},
			{Value: 10, Icon: " "},
			{Value: 0, Icon: " "},
		},
	},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    1 * time.Minute,
}

var brightness = core.Module{
	Name:   "brightness",
	Status: &providers.Brightness{},
	Icon: &core.ThresholdIcon{
		StatusConverterFn: core.NtoI,
		Thresholds: []core.Threshold{
			{Value: 50, Icon: "  "},
			{Value: 25, Icon: "  "},
			{Value: 0, Icon: "  "},
		},
	},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    1 * time.Minute,
}

var battery = core.Module{
	Name:   "battery",
	Status: &providers.Battery{Battery: "BAT1"},
	Icon: &providers.BatteryIconProvider{
		Battery:  "BAT1",
		Charging: "  ",
		ThresholdIcon: core.ThresholdIcon{
			StatusConverterFn: core.NtoI,
			Thresholds: []core.Threshold{
				{Value: 90, Icon: "  "},
				{Value: 60, Icon: "  "},
				{Value: 40, Icon: "  "},
				{Value: 10, Icon: "  "},
				{Value: 0, Icon: "  "},
			},
		},
	},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    15 * time.Second,
}

var layout = core.Module{
	Name:       "layout",
	Status:     &providers.Layout{},
	Icon:       &core.StaticIcon{Icon: "  "},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    1 * time.Hour,
}

var clock = core.Module{
	Name:       "clock",
	Status:     &providers.Clock{Layout: "2006-01-02 15:04"},
	Icon:       &core.StaticIcon{Icon: "  "},
	Gap:        core.DefaultGap,
	BlockOrder: core.IconFirst,
	Refresh:    30 * time.Second,
}
