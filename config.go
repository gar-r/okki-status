package main

import "bitbucket.org/dargzero/smart-status/modules"

type entry struct {
	module modules.Module
	format string
}

// 🔊

var config = []entry{
	{
		module: &modules.Volume{},
		format: " V: %s  ",
	},
	{
		module: &modules.Wifi{Device: "wlp59s0"},
		format: " N: %s  ",
	},
	{
		module: &modules.Battery{Battery: "BAT0", Charging: "(+)"},
		format: " B: %s  ",
	},
	{
		module: &modules.Clock{Layout: "2006-01-02 15:04"},
		format: " %s ",
	},
}
