package provider

import (
	"fmt"
	"okki-status/core"
	"time"

	"github.com/pbnjay/memory"
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
}

func (m *Memory) Run(ch chan<- core.Update, event <-chan core.Event) {
	ch <- m.getMemUsage()

	if m.Refresh == 0 {
		m.Refresh = 5000
	}

	ticker := time.NewTicker(time.Duration(m.Refresh) * time.Millisecond)
	for {
		<-ticker.C
		ch <- m.getMemUsage()
	}
}

func (m *Memory) getMemUsage() *MemoryUpdate {
	free := memory.FreeMemory()
	total := memory.TotalMemory()
	return &MemoryUpdate{
		p:     m,
		Free:  free,
		Total: total,
	}
}
