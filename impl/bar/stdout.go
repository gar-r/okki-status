package bar

import (
	"fmt"
	"okki-status/core"
)

type Stdout struct {
	Modules []*core.Module
}

func NewStdout(modules []*core.Module) *Stdout {
	s := &Stdout{
		Modules: modules,
	}
	for _, m := range s.Modules {
		m.Attach(s)
	}
	return s
}

func (s *Stdout) Update(e core.Event) {
	s.Render()
}

func (s *Stdout) Render() {
	for _, m := range s.Modules {
		fmt.Printf("%s\t", m.Status())
	}
}
