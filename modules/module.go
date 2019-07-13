package modules

import "fmt"

// Module is able to provide status for a single element on the status bar
type Module interface {
	Status() string
}

type Margin struct {
	Left, Right string
}

type Icon struct {
	Icon, Separator string
}

func (m *Margin) Format(s string) string {
	return fmt.Sprintf("%s%s%s", m.Left, s, m.Right)
}

func (i *Icon) Format(s string) string {
	return fmt.Sprintf("%s%s%s", i.Icon, i.Separator, s)
}
