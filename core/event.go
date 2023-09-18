package core

import (
	"encoding/json"
	"log"
	"os"

	sp "git.okki.hu/garric/swaybar-protocol"
)

// handleUpdates uses the event source to find the appropriate
// module and update the bar cache with its actual status.
// Once the cache is updated, the full bar needs to be rendered.
func (b *Bar) handleUpdates(update Update) {
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
		log.Println(err)
		return
	}
	if token != json.Delim('[') {
		log.Printf("unexpected start token: %s\n", token)
		return
	}

	// decode click events as they appear on stdin
	for dec.More() {
		ce := &sp.ClickEvent{}
		err := dec.Decode(ce)
		if err != nil {
			log.Println(err)
			return // something went wrong, give up
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
