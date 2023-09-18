package core

// Provider represents an information source that can
// push status updates to a module on the bar.
type Provider interface {
	// Run starts the provider.
	// A Provider is started exactly once by the Bar, and
	// then it is responsible for polling system updates,
	// or reacting to system events or click events.
	// The Provider has two channels: the first one can be used
	// to push updates to the Bar, the second one can be read
	// to receive external events that are addressed to the Module.
	Run(chan<- Update, <-chan Event)
}

// Update represents a status update.
// Providers should implement this interface to publish an update.
type Update interface {
	Source() Provider
	Text() string
}

type SimpleUpdate struct {
	P Provider
	T string
}

// SimpleUpdate is a basic implementation of the Update interface,
// containing only the provider and a text field.
func (s *SimpleUpdate) Source() Provider {
	return s.P
}

func (s *SimpleUpdate) Text() string {
	return s.T
}

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
