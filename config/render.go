package config

import (
	"time"

	"hu.okki.okki-status/core"
	"hu.okki.okki-status/renderer"
)

// global render interval
var Interval = 1 * time.Second

// rendering configuration
var R core.Renderer = &renderer.SwayBar{
	BlockCfg: []*renderer.SwayBarBlockConfig{
		{
			BlockName:      "battery",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "brightness",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "volume",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "ram",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "wifi",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "updates",
			SeparatorWidth: 25,
		},
		{
			BlockName:      "layout",
			SeparatorWidth: 25,
		},
	},
}
