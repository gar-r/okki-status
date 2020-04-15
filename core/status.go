package core

// StatusProvider supplies the status string
type StatusProvider interface {
	GetStatus() string
}
