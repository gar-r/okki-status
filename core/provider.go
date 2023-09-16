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
	// to receive mouse click events when the Module that belongs
	// to the Provider has been clicked.
	Run(chan<- *Update, <-chan *Click)
}

// Update represents a status update by one of the providers
// associated with a module on the bar.
type Update struct {
	Source      Provider
	Status      string
	StatusShort string
}

// Click represents a click event on one of the modules.
type Click struct {
	Name     string
	Instance string
	Button   int
	RelX     int
	RelY     int
}
