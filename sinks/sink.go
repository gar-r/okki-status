package sinks

type Sink interface {
	Accept(status string)
}
