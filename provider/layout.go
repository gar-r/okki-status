package provider

import (
	"encoding/json"
	"io"
	"okki-status/core"
	"os/exec"
	"strings"
)

const (
	inputTypeKeyboard = "keyboard"
	typeLayout        = "xkb_layout"
)

// Layout provides keyboard layout information under swaywm
type Layout struct {
	Identifier string `yaml:"keyboard_identifier"`
	Shorten    bool   `yaml:"shorten"`
}

func (l *Layout) Run(ch chan<- core.Update, event <-chan core.Event) {
	layout, err := l.currentLayout()
	if err != nil {
		ch <- &core.ErrorUpdate{P: l}
		return
	}
	l.sendLayoutUpdate(ch, layout)
	go l.listenForClickEvents(event)
	l.listenForUpdates(ch)
}

func (l *Layout) currentLayout() (string, error) {
	output, err := exec.Command("swaymsg", "-r", "-t", "get_inputs").Output()
	if err != nil {
		return "", err
	}
	var inputs []Input
	err = json.Unmarshal(output, &inputs)
	if err != nil {
		return "", err
	}
	for _, input := range inputs {
		if input.Identifier == l.Identifier && input.Type == inputTypeKeyboard {
			return input.Layout, nil
		}
	}
	return "", err
}

func (l *Layout) listenForUpdates(ch chan<- core.Update) {
	cmd := exec.Command("swaymsg", "-r", "-m", "-t", "subscribe", "[\"input\"]")
	r, err := cmd.StdoutPipe()
	if err != nil {
		ch <- &core.ErrorUpdate{P: l}
		return
	}
	err = cmd.Start()
	if err != nil {
		ch <- &core.ErrorUpdate{P: l}
	}
	dec := json.NewDecoder(r)
	for {
		change := &InputChange{}
		err = dec.Decode(change)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue // decoding failed, nothing to do
		}
		if change.Change == typeLayout &&
			change.Input.Type == inputTypeKeyboard &&
			change.Input.Identifier == l.Identifier {
			l.sendLayoutUpdate(ch, change.Input.Layout)
		}
	}
}

func (l *Layout) listenForClickEvents(event <-chan core.Event) {
	for {
		<-event
		cmd := exec.Command("swaymsg", "input", "type:keyboard", "xkb_switch_layout", "next")
		_ = cmd.Run()
	}
}

func (l *Layout) sendLayoutUpdate(ch chan<- core.Update, layout string) {
	var text string
	if l.Shorten {
		text = shorten(layout)
	} else {
		text = layout
	}
	ch <- &core.SimpleUpdate{
		P: l,
		T: text,
	}
}

func shorten(layout string) string {
	layoutLower := strings.ToLower(layout)
	if len(layoutLower) < 2 {
		return layoutLower
	}
	return layoutLower[:2]
}

// Input represents a sway input
type Input struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
	Layout     string `json:"xkb_active_layout_name"`
}

// InputChange represents an input change event from swaywm
type InputChange struct {
	Input  *Input `json:"input"`
	Change string `json:"change"`
}
