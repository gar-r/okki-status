package core

import (
	"os"

	sp "git.okki.hu/garric/swaybar-protocol"
)

type Bar struct {
	eventbus chan *Event
	Modules  []*Module `yaml:"modules"`
	cache    []*sp.Body
}

func (b *Bar) Start() {
	b.eventbus = make(chan *Event)
	b.cache = make([]*sp.Body, len(b.Modules))
	for _, m := range b.Modules {
		go m.Run(b.eventbus)
	}
	for {
		b.handle(<-b.eventbus)
	}
}

func (b *Bar) Render() {
	sp.Output(os.Stdout, b.cache)
}

func (b *Bar) handle(e *Event) {
	for i, module := range b.Modules {
		if module.Provider == e.Source {
			b.cache[i] = module.Render(e.Status)
		}
	}
	b.Render()
}

type Event struct {
	Source Provider
	Status string
}
