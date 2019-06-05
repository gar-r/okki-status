package sinks

// Sink is a target for status information
type Sink interface {
	Accept(status string)
}
