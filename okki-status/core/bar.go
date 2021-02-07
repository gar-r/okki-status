package core

import (
	"log"
	"strings"

	"hu.okki.okki-status/output"
)

// Bar is composed of an arbitrary number of modules
type Bar struct {
	modules []Module
	values  map[Module]string
}

func NewBar(modules []Module) *Bar {
	return &Bar{
		modules: modules,
		values:  make(map[Module]string, len(modules)),
	}
}

func (b *Bar) Render(sink output.Sink) {
	output := strings.Builder{}
	for _, module := range b.modules {
		output.WriteString(b.fetch(module))
	}
	sink.Accept(output.String())
}

func (b *Bar) Invalidate(module Module) {
	log.Printf("invalidating %s", module.Name)
	b.values[module] = module.Info()
}

func (b *Bar) fetch(module Module) (value string) {
	value, present := b.values[module]
	if !present {
		value = module.Info()
		b.values[module] = value
	}
	return
}
