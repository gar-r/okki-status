package core

import (
	"encoding/json"
	"errors"
	"os"

	sp "git.okki.hu/garric/swaybar-protocol"
)

// Bar represents the swaybar
type Bar struct {
	errors  chan error
	updates chan Update
	Modules []*Module `yaml:"modules"`
	cache   []*sp.Body
}

// Start the swaybar
func (b *Bar) Start() error {
	b.errors = make(chan error)
	b.updates = make(chan Update)
	b.cache = make([]*sp.Body, len(b.Modules))
	if err := b.renderHeader(); err != nil {
		return err
	}
	go b.processClicks()
	go b.startDbusServer()
	b.startModules()
	for {
		select {
		case err := <-b.errors:
			return err
		case update := <-b.updates:
			b.handleUpdate(update)
		}
	}
}

// renderHeader uses swaybar-protocol to output a header
func (b *Bar) renderHeader() error {
	return sp.Init(os.Stdout, &sp.Header{
		Version:     1,
		ClickEvents: true,
	})
}

// render uses swaybar-protocol to output the modules from cache
func (b *Bar) render() {
	err := sp.Output(os.Stdout, b.cache)
	if err != nil {
		b.errors <- err
	}
}

// startModules starts each module in a separate go func,
// and initializes their click event channel
func (b *Bar) startModules() {
	for _, m := range b.Modules {
		m.events = make(chan Event)
		go m.Run(b.updates, m.events)
	}
}

// handleUpdate uses the event source to find the appropriate
// module and update the bar cache with its actual status.
// Once the cache is updated, the full bar needs to be rendered.
func (b *Bar) handleUpdate(update Update) {
	for i, module := range b.Modules {
		if module.Provider == update.Source() {
			b.cache[i] = module.Render(update)
		}
	}
	b.render()
}

// processClicks reads click events sent to stdin by swaybar-protocol.
// When an event is read, it identifies the clicked module, and
// publishes on its Click channel.
// This will result in a Click event only being sent to the Module
// that was clicked.
func (b *Bar) processClicks() {
	dec := json.NewDecoder(os.Stdin)

	// read begin array token ('[')
	token, err := dec.Token()
	if err != nil {
		b.errors <- err
	}
	if token != json.Delim('[') {
		b.errors <- errors.New("unexpected start delimiter on stdin")
	}

	// decode click events as they appear on stdin
	for dec.More() {
		ce := &sp.ClickEvent{}
		err := dec.Decode(ce)
		if err != nil {
			b.errors <- err
		}
		for _, module := range b.Modules {
			if module.Name == ce.Name {
				module.events <- &Click{
					Name:     ce.Name,
					Instance: ce.Instance,
					Button:   ce.Button,
					RelX:     ce.RelativeX,
					RelY:     ce.RelativeY,
				}
			}
		}
	}
}
