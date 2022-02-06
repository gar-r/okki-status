package module

import (
	"encoding/json"
	"os/exec"

	"hu.okki.okki-status/core"
)

// KeyboardSw provides keyboard layout information under Sway
type KeyboardSw struct {
	// keyboard identifier
	Identifier string
}

// Input represents a sway input
type Input struct {
	Identifier string `json:"identifier"`
	Layout     string `json:"xkb_active_layout_name"`
}

// Status returns the keyboard layout string
func (k *KeyboardSw) Status() string {
	output, err := exec.Command("swaymsg", "-t", "get_inputs").Output()
	if err != nil {
		return core.StatusError
	}
	var inputs []Input
	err = json.Unmarshal(output, &inputs)
	if err != nil {
		return core.StatusError
	}
	for _, input := range inputs {
		if input.Identifier == k.Identifier {
			if input.Layout != "" && len(input.Layout) > 1 {
				return input.Layout[:2]
			}
		}
	}
	return core.StatusUnknown
}
