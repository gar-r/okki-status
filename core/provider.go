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

// ErrorUpdate is typically sent when the module state is in error.
type ErrorUpdate struct {
	P Provider
}

func (e *ErrorUpdate) Source() Provider {
	return e.P
}

func (e *ErrorUpdate) Text() string {
	return "?"
}
