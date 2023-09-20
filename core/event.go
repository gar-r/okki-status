package core

// Event represents an external event that a provider can
// choose to react to.
type Event interface {
	ModuleName() string
}

// Click represents a click event on one of the modules.
type Click struct {
	Name     string
	Instance string
	Button   int
	RelX     int
	RelY     int
}

func (c *Click) ModuleName() string {
	return c.Name
}

// Refresh represents a refresh request for a module.
type Refresh struct {
	Name string
}

func (r *Refresh) ModuleName() string {
	return r.Name
}
