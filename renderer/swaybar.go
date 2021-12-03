package renderer

import (
	"encoding/json"
	"fmt"
	"log"

	"hu.okki.okki-status/core"
)

// SwayBar renders using the sway-bar json protocol.
// The protocol comprises a header, that is a JSON object,
// followed by a newline (0x0A), followed by an infinite JSON
// array that represents the information to display.
// All communication is done by writing the status line to
// standard output and reading events from standard input.
type SwayBar struct {
	BlockCfg []*SwayBarBlockConfig
	started  bool
}

func (s *SwayBar) Render(bar core.Bar) {
	if !s.started {
		swayout(marshal(s.makeHeader()))
		swayout("\n")
		swayout("[")
		s.started = true
	}
	swayout(marshal(s.makeSwayBar(bar)))
	swayout(",")
}

func (s *SwayBar) getConfig(block *core.Block) *SwayBarBlockConfig {
	for _, cfg := range s.BlockCfg {
		if cfg.BlockName == block.Name {
			return cfg
		}
	}
	return &SwayBarBlockConfig{BlockName: block.Name}
}

func (s *SwayBar) makeHeader() *SwayBarHeader {
	return &SwayBarHeader{Version: "1"}
}

func (s *SwayBar) makeSwayBar(bar core.Bar) []*SwayBarBlock {
	arr := make([]*SwayBarBlock, 0)
	for _, block := range bar {
		cfg := s.getConfig(block)
		el := s.makeSwayBarBlock(block, cfg)
		arr = append(arr, el)
	}
	return arr
}

func (s *SwayBar) makeSwayBarBlock(block *core.Block, cfg *SwayBarBlockConfig) *SwayBarBlock {
	return &SwayBarBlock{
		Name:           cfg.BlockName,
		FullText:       block.Status(),
		Color:          cfg.TextColor,
		Background:     cfg.BackgroundColor,
		Align:          string(cfg.Alignment),
		Separator:      cfg.Separator,
		SeparatorWidth: cfg.SeparatorWidth,
	}
}

// SwayBarBlockConfig contains block specfic config.
// Attributes are only applied when rendering a block with a matching name.
type SwayBarBlockConfig struct {
	BlockName       string    // the name of the Block
	TextColor       string    // the text color (using #RRGGBB[AA] notation)
	BackgroundColor string    // the background color (using #RRGGBB[AA] notation)
	Alignment       SwayAlign // the alignment (left|right|center)
	Separator       bool      // do not render separator block after this block
	SeparatorWidth  int       // width of the separator block
}

type SwayAlign string

const (
	Left   = "left"
	Right  = "right"
	Center = "center"
)

// SwayBarHeader is used to output protocol header json data
type SwayBarHeader struct {
	Version string `json:"version"`
}

// SwayBarBlock is used to output protocol body json data
type SwayBarBlock struct {
	Name           string `json:"name"`
	FullText       string `json:"full_text"`
	Color          string `json:"color,omitempty"`
	Background     string `json:"background,omitempty"`
	Align          string `json:"align,omitempty"`
	Separator      bool   `json:"separator"`
	SeparatorWidth int    `json:"separator_block_width,omitempty"`
}

func marshal(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Println("error while serializing header", err.Error())
	}
	return string(b)
}

var swayout = func(s string) {
	fmt.Print(s)
}
