package core

import (
	"os"

	sp "git.okki.hu/garric/swaybar-protocol"
)

// Bar represents the swaybar
type Bar struct {
	updch   chan *Update
	Modules []*Module `yaml:"modules"`
	cache   []*sp.Body
}

// Start the swaybar
func (b *Bar) Start() {
	b.updch = make(chan *Update)
	b.cache = make([]*sp.Body, len(b.Modules))
	b.renderHeader()
	go b.processClicks()
	b.startModules()
	for {
		b.handleUpdates(<-b.updch)
	}
}

// renderHeader uses swaybar-protocol to output a header
func (b *Bar) renderHeader() {
	sp.Init(os.Stdout, &sp.Header{
		Version:     1,
		ClickEvents: true,
	})
}

// render uses swaybar-protocol to output the modules from cache
func (b *Bar) render() {
	sp.Output(os.Stdout, b.cache)
}

// startModules starts each module in a separate go func,
// and initializes their click event channel
func (b *Bar) startModules() {
	for _, m := range b.Modules {
		m.clkch = make(chan *Click)
		go m.Run(b.updch, m.clkch)
	}
}
