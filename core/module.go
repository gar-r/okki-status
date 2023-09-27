package core

import (
	"text/template"

	sp "git.okki.hu/garric/swaybar-protocol"
)

// Module represents a block on the bar.
// It encapsulates the Appearance, a Provider and a channel
// through which the module can receive events.
// Each module is uniquely identitied by its Name.
type Module struct {
	Provider     `yaml:"-"`
	ProviderConf map[string]interface{} `yaml:"provider"`
	Appearance   *Appearance            `yaml:"appearance"`
	events       chan Event
	Name         string     `yaml:"name"`
	Variants     []*Variant `yaml:"variants"`
}

// Render transforms the current state of the module into a
// swaybar-protocol body object, which can be directly sent to
// swaybar for output.
// The module will use its current state when generating
// the body object, including the Update sent by the Provider,
// the Appearance, and Appearance Variants.
func (m *Module) Render(update Update) *sp.Body {
	a := m.getAppearance(update)
	body := &sp.Body{
		Name:     m.Name,
		MinWidth: a.MinWidth,
		Align:    a.Align,
		Urgent:   a.Urgent,
	}

	if a.Color != nil {
		body.Color = a.Color.Foreground
		body.Background = a.Color.Background
	}

	if a.Border != nil {
		body.Border = a.Border.Color
		body.BorderTop = a.Border.Top
		body.BorderBottom = a.Border.Bottom
		body.BorderLeft = a.Border.Left
		body.BorderRight = a.Border.Right
	}

	if a.Separator != nil {
		body.Separator = a.Separator.Enabled
		body.SeparatorBlockWidth = a.Separator.BlockWidth
	}

	// add full and short status text
	body.FullText = a.ExecuteFormat(update)
	body.ShortText = a.ExecuteFormatShort(update)

	return body
}

// calculate the module appearance based on the configured defaults
// and the first matching Variant (if any)
func (m *Module) getAppearance(update Update) *Appearance {
	base := m.Appearance
	var variant *Appearance
	for _, v := range m.Variants {
		if v.Match(update) {
			variant = v.Appearance
			break
		}
	}
	if variant == nil {
		return base
	}
	appearance := &Appearance{}
	appearance.Format = override(base.Format, variant.Format)
	appearance.FormatShort = override(base.Format, variant.FormatShort)
	appearance.formatCompiled = override(base.formatCompiled, variant.formatCompiled)
	appearance.formatShortCompiled = override(base.formatShortCompiled, variant.formatShortCompiled)
	appearance.MinWidth = override(base.MinWidth, variant.MinWidth)
	appearance.Align = override(base.Align, variant.Align)
	appearance.Urgent = override(base.Urgent, variant.Urgent)
	if variant.Color == nil {
		appearance.Color = base.Color
	} else {
		if base.Color == nil {
			appearance.Color = variant.Color
		} else {
			appearance.Color = &Color{
				Foreground: override(base.Color.Foreground, variant.Color.Foreground),
				Background: override(base.Color.Background, variant.Color.Background),
			}
		}
	}
	if variant.Border == nil {
		appearance.Border = base.Border
	} else {
		if base.Border == nil {
			appearance.Border = variant.Border
		} else {
			appearance.Border = &Border{
				Color:  override(base.Border.Color, variant.Border.Color),
				Top:    override(base.Border.Top, variant.Border.Top),
				Bottom: override(base.Border.Bottom, variant.Border.Bottom),
				Left:   override(base.Border.Left, variant.Border.Left),
				Right:  override(base.Border.Right, variant.Border.Right),
			}
		}
	}
	if variant.Separator == nil {
		appearance.Separator = base.Separator
	} else {
		if base.Separator == nil {
			appearance.Separator = variant.Separator
		} else {
			appearance.Separator = &Separator{
				Enabled:    override(base.Separator.Enabled, variant.Separator.Enabled),
				BlockWidth: override(base.Separator.BlockWidth, variant.Separator.BlockWidth),
			}
		}
	}
	return appearance
}

// pick between the base and override value
func override[T string | int | bool | *template.Template](base, override T) T {
	var zero T
	if override == zero {
		return base
	}
	return override
}
