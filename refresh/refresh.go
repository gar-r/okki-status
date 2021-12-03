package refresh

import (
	"hu.okki.okki-status/config"
	"hu.okki.okki-status/core"
)

func find(name string) core.Module {
	for _, b := range config.B {
		if name == b.Name {
			return b.Module
		}
	}
	return nil
}

func refresh(m core.Module) {
	invalidate(m)
	config.R.Render(config.B)
}

func invalidate(m core.Module) {
	// if it is caching, invalidate the cached value
	cm, ok := m.(*core.CachingModule)
	if ok {
		cm.Invalidate()
	}
}
