package module

// Text provides static text
type Text struct {
	T string
}

// Status simply returns the value of S
func (m *Text) Status() string {
	return m.T
}
