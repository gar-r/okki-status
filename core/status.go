package core

type StatusProvider interface {
	GetStatus() string
}
