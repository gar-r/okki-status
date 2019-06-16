package main

import "bitbucket.org/dargzero/smart-status/modules"

type entry struct {
	module modules.Module
	format string
}

var config = []entry{
	{
		module: &modules.Wifi{Device: "wlp59s0"},
		format: "    %s   ",
	},
	{
		module: &modules.RAM{},
		format: "    %s   ",
	},
	{
		module: &modules.Volume{},
		format: "    %s   ",
	},
	{
		module: &modules.Brightness{},
		format: "    %s   ",
	},
	{
		module: &modules.Battery{Battery: "BAT0", Charging: " ", Discharging: " "},
		format: "   %s   ",
	},
	{
		module: &modules.Clock{Layout: "2006-01-02 15:04"},
		format: "    %s  ",
	},
}
