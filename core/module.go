package core

type Provider interface {
	Status() string
}

type Module struct {
	Provider     `yaml:"-"`
	Name         string                 `yaml:"name"`
	ProviderConf map[string]interface{} `yaml:"provider"`
	observers    []Observer
}

func (m *Module) Attach(o Observer) {
	m.observers = append(m.observers, o)
}

func (m *Module) Notify() {
	for _, o := range m.observers {
		e := Event{m.Status()}
		o.Update(e)
	}
}
