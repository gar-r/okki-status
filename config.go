package main

import "bitbucket.org/dargzero/smart-status/modules"

var config = []modules.Module{
	&modules.Wifi{
		Device: "wlp1s0",
	},
	&modules.RAM{
		Margin: defaultMargin,
		Icon:   iconWithGap(""),
	},
	&modules.Volume{
		Margin: defaultMargin,
		Icon:   iconWithGap(""),
	},
	&modules.Brightness{
		Margin: defaultMargin,
		Icon:   iconWithGap(""),
	},
	&modules.Battery{
		Battery: "BAT0",
		StatusMap: map[string]string{
			"Charging":    "",
			"Discharging": "",
			"Full":        "",
		},
		Margin: defaultMargin,
	},
	&modules.Clock{
		Layout: "2006-01-02 15:04",
		Margin: defaultMargin,
		Icon:   iconWithGap(""),
	},
}

var defaultMargin = modules.Margin{
	Left:  "   ",
	Right: "   ",
}

func iconWithGap(icon string) modules.Icon {
	return modules.Icon{
		Icon:      icon,
		Separator: " ",
	}
}
