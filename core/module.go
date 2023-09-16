package core

import (
	"fmt"

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
// the body object, including the Update sent by the Provider,
// the Appearance, and Appearance Variants.
func (m *Module) Render(status *Update) *sp.Body {
	// make body with configured settings
	body := &sp.Body{
		Color:               m.Appearance.Color.Foreground,
		Background:          m.Appearance.Color.Background,
		Border:              m.Appearance.Color.Border,
		BorderTop:           m.Appearance.Border.Top,
		BorderBottom:        m.Appearance.Border.Bottom,
		BorderLeft:          m.Appearance.Border.Left,
		BorderRight:         m.Appearance.Border.Right,
		MinWidth:            m.Appearance.MinWidth,
		Align:               m.Appearance.Align,
		Name:                m.Name,
		Urgent:              m.Appearance.Urgent,
		Separator:           m.Appearance.Separator.Enabled,
		SeparatorBlockWidth: m.Appearance.Separator.BlockWidth,
	}

	// add full and short status text
	body.FullText = fmt.Sprintf(m.Appearance.Format, status.Status)
	body.ShortText = fmt.Sprintf(m.Appearance.FormatShort, status.StatusShort)
	return body

	// TODO: implement appearance variants
}
