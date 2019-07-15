package main

import (
	"bitbucket.org/dargzero/smart-status/core"
	"bitbucket.org/dargzero/smart-status/providers"
	"strconv"
	"strings"
)

var config = Config{
	{
		StatusProvider: &providers.Wifi{Device: "wlp1s0"},
		IconProvider:   &core.StaticIcon{Icon: " "},
		Gap:            defaultGap,
		BlockOrder:     core.IconFirst,
	},
	{
		StatusProvider: &providers.RAM{},
		IconProvider:   &core.StaticIcon{Icon: " "},
		Gap:            defaultGap,
		BlockOrder:     core.IconFirst,
	},
	{
		StatusProvider: &providers.Volume{},
		IconProvider: &core.ThresholdIcon{
			StatusConverterFn: valToPercent,
			Thresholds: []core.Threshold{
				{Value: 75, Icon: " "},
				{Value: 0, Icon: " "},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
	},
	{
		StatusProvider: &providers.Brightness{},
		IconProvider: &core.ThresholdIcon{
			StatusConverterFn: valToPercent,
			Thresholds: []core.Threshold{
				{Value: 50, Icon: " "},
				{Value: 25, Icon: " "},
				{Value: 0, Icon: " "},
			},
		},
		Gap:        defaultGap,
		BlockOrder: core.IconFirst,
	},
	{
		StatusProvider: &providers.Battery{Battery: "BAT0"},
		IconProvider: &providers.BatteryIconProvider{
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
	},
	{
		StatusProvider: &providers.Clock{Layout: "2006-01-02 15:04"},
		IconProvider:   &core.StaticIcon{Icon: " "},
		Gap:            defaultGap,
		BlockOrder:     core.IconFirst,
	},
}

type Config []core.Block

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
