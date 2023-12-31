package provider

import (
	"fmt"
	"okki-status/core"
	"time"

	sigar "github.com/cloudfoundry/gosigar"
)

type MemoryUpdate struct {
	p     core.Provider
	Free  uint64
	Total uint64
}

func (m *MemoryUpdate) Source() core.Provider {
	return m.p
}

func (m *MemoryUpdate) Text() string {
	return fmt.Sprintf("%.0f%%", float64(m.Used())/float64(m.Total)*100)
}

func (m *MemoryUpdate) Used() uint64 {
	return m.Total - m.Free
}

type Memory struct {
	Refresh int
	mem     *sigar.Mem
}

func (m *Memory) Run(ch chan<- core.Update, event <-chan core.Event) {
	m.mem = &sigar.Mem{}
	if m.Refresh == 0 {
		m.Refresh = 5000
	}

	ch <- m.getMemUsage()
	ticker := time.NewTicker(time.Duration(m.Refresh) * time.Millisecond)
	for {
		<-ticker.C
		ch <- m.getMemUsage()
	}
}

func (m *Memory) getMemUsage() *MemoryUpdate {
	m.mem.Get()
	free := m.mem.Free
	total := m.mem.Total
	return &MemoryUpdate{
		p:     m,
		Free:  free,
		Total: total,
	}
}
