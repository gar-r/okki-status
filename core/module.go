package core

import (
	sp "git.okki.hu/garric/swaybar-protocol"
)

type Provider interface {
	Run(chan<- *Event)
}
type Module struct {
	Provider     `yaml:"-"`
	ProviderConf map[string]interface{} `yaml:"provider"`
	Name         string                 `yaml:"name"`
	Appearance   *Appearance            `yaml:"appearance"`
	Variants     []*Variant             `yaml:"variants"`
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
func (m *Module) Render(status string) *sp.Body {
	return &sp.Body{
		Name: m.Name,
		// TODO: set appearance, variant, etc
	}
}
