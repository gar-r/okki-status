package main

import (
	"bitbucket.org/dargzero/smart-status/core"
	"bitbucket.org/dargzero/smart-status/providers"
	"strconv"
	"strings"
	"time"
)

var addr = ":12650"

var config = Config{
	{
		Name:       "wifi",
		Status:     &providers.Wifi{Device: "wlp1s0"},
		Icon:       &core.StaticIcon{Icon: " "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    5 * time.Second,
	},
	{
		Name:       "ram",
		Status:     &providers.RAM{},
		Icon:       &core.StaticIcon{Icon: " "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    3 * time.Second,
	},
	{
		Name:   "volume",
		Status: &providers.Volume{},
		Icon: &core.ThresholdIcon{
			StatusConverterFn: valToPercent,
			Thresholds: []core.Threshold{
				{Value: 75, Icon: " "},
				{Value: 0, Icon: " "},
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
			StatusConverterFn: valToPercent,
			Thresholds: []core.Threshold{
				{Value: 50, Icon: " "},
				{Value: 25, Icon: " "},
				{Value: 0, Icon: " "},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Minute,
	},
	{
		Name:   "battery",
		Status: &providers.Battery{Battery: "BAT0"},
		Icon: &providers.BatteryIconProvider{
			Battery:  "BAT0",
			Charging: " ",
			ThresholdIcon: core.ThresholdIcon{
				StatusConverterFn: valToPercent,
				Thresholds: []core.Threshold{
					{Value: 90, Icon: " "},
					{Value: 60, Icon: " "},
					{Value: 40, Icon: " "},
					{Value: 10, Icon: " "},
					{Value: 0, Icon: " "},
				},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    5 * time.Second,
	},
	{
		Name:       "clock",
		Status:     &providers.Clock{Layout: "2006-01-02 15:04"},
		Icon:       &core.StaticIcon{Icon: " "},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
		Refresh:    1 * time.Minute,
	},
}

type Config []core.Module

var defaultGap = core.Gap{
	Before: "   ",
	After:  "   ",
}

func valToPercent(value string) int {
	percent, err := strconv.Atoi(strings.Replace(value, "%", "", 1))
	if err != nil {
		return -1
	}
	return percent
}
