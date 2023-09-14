package core

type Provider interface {
	Status() string
}

type Module struct {
	Provider     `yaml:"-"`
	Name         string                 `yaml:"name"`
	Appearance   *Appearance            `yaml:"appearance"`
	Alternates   []*Alternate           `yaml:"alternates"`
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

// FullText:            "fullText",
// ShortText:           "shortText",
// Color:               "#ccccccff",
// Background:          "#111111ff",
// Border:              "#222222ff",
// BorderTop:           1,
// BorderBottom:        1,
// BorderLeft:          1,
// BorderRight:         1,
// MinWidth:            100,
// Align:               AlignCenter,
// Name:                "name",
// Instance:            "instance",
// Urgent:              true,
// Separator:           true,
// SeparatorBlockWidth: 5,
// Markup:              MarkupNone,
//
