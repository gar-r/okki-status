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

// Render transforms the current state of the module into a
// swaybar-protocol body object, which can be directly sent to
// swaybar for output.
// The module will use its current state when generating
// the body object, including the Update sent by the Provider,
// the Appearance, and Appearance Variants.
func (m *Module) Render(update *Update) *sp.Body {
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
		body.Border = a.Color.Border
	}

	if a.Border != nil {
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
	body.FullText = fmt.Sprintf(a.Format, update.Status)
	if m.Appearance.FormatShort != "" {
		body.ShortText = fmt.Sprintf(a.FormatShort, update.StatusShort)
	}

	return body
}

// calculate the module appearance based on the configured defaults
// and the first matching Variant (if any)
func (m *Module) getAppearance(update *Update) *Appearance {
	base := m.Appearance
	var variant *Appearance
	for _, v := range m.Variants {
		if v.Match(update.Status) || v.MatchShort(update.StatusShort) {
			variant = v.Appearance
			break
		}
	}
	if variant == nil {
		return base
	}
	appearance := &Appearance{}
	appearance.Format = overrideStr(base.Format, variant.Format)
	appearance.FormatShort = overrideStr(base.Format, variant.FormatShort)
	appearance.MinWidth = overrideInt(base.MinWidth, variant.MinWidth)
	appearance.Align = overrideStr(base.Align, variant.Align)
	appearance.Urgent = overrideBool(base.Urgent, variant.Urgent)
	if variant.Color == nil {
		appearance.Color = base.Color
	} else {
		if base.Color == nil {
			appearance.Color = variant.Color
		} else {
			appearance.Color = &Color{
				Foreground: overrideStr(base.Color.Foreground, variant.Color.Foreground),
				Background: overrideStr(base.Color.Background, variant.Color.Background),
				Border:     overrideStr(base.Color.Border, variant.Color.Border),
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
				Top:    overrideInt(base.Border.Top, variant.Border.Top),
				Bottom: overrideInt(base.Border.Bottom, variant.Border.Bottom),
				Left:   overrideInt(base.Border.Left, variant.Border.Left),
				Right:  overrideInt(base.Border.Right, variant.Border.Right),
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
				Enabled:    overrideBool(base.Separator.Enabled, variant.Separator.Enabled),
				BlockWidth: overrideInt(base.Separator.BlockWidth, variant.Separator.BlockWidth),
			}
		}
	}
	return appearance
}

func overrideStr(base, override string) string {
	if override == "" {
		return base
	}
	return override
}

func overrideInt(base, override int) int {
	return 0
}

func overrideBool(base, override bool) bool {
	if override {
		return true
	}
	return base
}
