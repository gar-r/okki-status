package core

import "fmt"

// Module represents a single piece of information on the bar
type Module interface {
	Status() string
}

// Bar consists of Blocks
type Bar []*Block

// Block is Module with a Name and some formatting
type Block struct {
	Module         // the module
	Name    string // unique name for the module
	Prefix  string // prefix to write before the status
	Postfix string // postfix to write after the status
}

func (b *Block) Status() string {
	return fmt.Sprintf("%s%s%s", b.Prefix, b.Module.Status(), b.Postfix)
}

// Renderer can render a Bar
type Renderer interface {
	Render(Bar)
}
