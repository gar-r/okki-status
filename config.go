package main

import (
	"strconv"
	"strings"
	"time"

	"bitbucket.org/dargzero/okki-status/core"
	"bitbucket.org/dargzero/okki-status/providers"
)

var addr = ":12650"

var config = Config{
	{
		Name:       "updates",
		Status:     &providers.Updates{Command: "/usr/bin/checkupdates"},
		Icon:       &core.StaticIcon{Icon: "⟑  "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    10 * time.Minute,
	},
	{
		Name:       "wiFi",
		Status:     &providers.WiFi{Device: "wlp0s20f3"},
		Icon:       &core.StaticIcon{Icon: "  "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    5 * time.Second,
	},
	{
		Name:       "ram",
		Status:     &providers.RAM{},
		Icon:       &core.StaticIcon{Icon: "  "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    15 * time.Second,
	},
	{
		Name:   "volume",
		Status: &providers.Volume{},
		Icon: &core.ThresholdIcon{
			StatusConverterFn: NtoI,
			Thresholds: []core.Threshold{
				{Value: 75, Icon: "  "},
				{Value: 0, Icon: "  "},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Minute,
	},
	{
		Name:   "brightness",
		Status: &providers.Brightness{},
		Icon: &core.ThresholdIcon{
			StatusConverterFn: NtoI,
			Thresholds: []core.Threshold{
				{Value: 50, Icon: "  "},
				{Value: 25, Icon: "  "},
				{Value: 0, Icon: "  "},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Minute,
	},
	{
		Name:   "battery",
		Status: &providers.Battery{Battery: "BAT1"},
		Icon: &providers.BatteryIconProvider{
			Battery:  "BAT1",
			Charging: "  ",
			ThresholdIcon: core.ThresholdIcon{
				StatusConverterFn: NtoI,
				Thresholds: []core.Threshold{
					{Value: 90, Icon: "  "},
					{Value: 60, Icon: "  "},
					{Value: 40, Icon: "  "},
					{Value: 10, Icon: "  "},
					{Value: 0, Icon: "  "},
				},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    15 * time.Second,
	},
	{
		Name:       "layout",
		Status:     &providers.Layout{},
		Icon:       &core.StaticIcon{Icon: "  "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Hour,
	},
	{
		Name:       "clock",
		Status:     &providers.Clock{Layout: "2006-01-02 15:04"},
		Icon:       &core.StaticIcon{Icon: "  "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Minute,
	},
}

// Config is composed of an arbitrary number of modules
type Config []core.Module

var defaultGap = core.Gap{
	Before: "   ",
	After:  "   ",
}

// NtoI converts a number-string to int - the number string can also
// be a percentage. For example "45%"" will become the int value of 45.
func NtoI(value string) int {
	n, err := strconv.Atoi(strings.Replace(value, "%", "", 1))
	if err != nil {
		return -1
	}
	return n
}
