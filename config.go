package main

import "bitbucket.org/dargzero/smart-status/modules"

type entry struct {
	module modules.Module
	format string
}

// 🔊

var config = []entry{
	entry{
		module: &modules.Wifi{Device: "wlp59s0"},
		format: " 🔗 %s  ",
	},
	entry{
		module: &modules.Battery{Battery: "BAT0", Charging: "⚡"},
		format: " 🔋 %s  ",
	},
	entry{
		module: &modules.Clock{Layout: "2006-01-02 15:04"},
		format: " 🕒 %s ",
	},
}
