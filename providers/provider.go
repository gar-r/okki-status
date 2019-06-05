package providers

// Provider provides data for a single piece of status information
type Provider interface {
	GetData() string
}
