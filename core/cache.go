package core

import "time"

// CachingModule is a wrapper around a single Module providing caching
type CachingModule struct {
	Module               // the wrapped module
	value   string       // the cached value
	expired bool         // is the cached value expired
	ticker  *time.Ticker // expiry ticker
}

func NewCachingModule(m Module, expiry time.Duration) *CachingModule {
	cm := &CachingModule{Module: m}
	cm.refresh()
	cm.ticker = time.NewTicker(expiry)
	go func() {
		for range cm.ticker.C {
			cm.Invalidate()
		}
	}()
	return cm
}

func (m *CachingModule) Status() string {
	if m.expired {
		m.refresh()
	}
	return m.value
}

func (m *CachingModule) Invalidate() {
	m.expired = true
}

func (m *CachingModule) refresh() {
	m.value = m.Module.Status()
	m.expired = false
}
