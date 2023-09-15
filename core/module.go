package core

import (
	sp "git.okki.hu/garric/swaybar-protocol"
)

// Module represents a block on the bar.
// It encapsulates the Appearance, a Provider and a channel
// through which the module can receive mouse click events.
// Each module is uniquely identitied by its Name.
type Module struct {
	Provider     `yaml:"-"`
	ProviderConf map[string]interface{} `yaml:"provider"`
	Appearance   *Appearance            `yaml:"appearance"`
	clkch        chan *Click
	Name         string     `yaml:"name"`
	Variants     []*Variant `yaml:"variants"`
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

// Render transforms the current state of the module into a
// swaybar-protocol body object, which can be directly sent to
// swaybar for output.
// The module will use its current state when generating
// the body object, including the status sent by the Provider,
// the Appearance, and Appearance Variants, while also setting
// sensible defaults for settings that were not specified through
// the configuration file.
func (m *Module) Render(status string) *sp.Body {
	return &sp.Body{
		Name:     m.Name,
		FullText: status,
		Align:    sp.AlignRight,
		// TODO: implement format text
		// TODO: set appearance
		// TODO: implement appearance variants
	}
}
