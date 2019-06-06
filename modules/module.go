package modules

// Module is able to provide status for a single element on the status bar
type Module interface {
	Status() string
}
