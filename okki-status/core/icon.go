package core

// IconProvider is able to supply an icon
type IconProvider interface {
	GetIcon(status string) string
}

// StaticIcon provides a simple static icon
type StaticIcon struct {
	Icon string
}

// GetIcon returns the icon
func (i *StaticIcon) GetIcon(status string) string {
	return i.Icon
}

// ThresholdIcon provides different icons based on predefined thresholds
type ThresholdIcon struct {
	// Thresholds contains the value thresholds in descending order
	Thresholds []Threshold

	// StatusConverterFn converts the status string to a numeric
	// value which is compared to the threshold
	StatusConverterFn func(string) int
}

// GetIcon returns the icon
func (g *ThresholdIcon) GetIcon(status string) string {
	value := g.StatusConverterFn(status)
	for _, r := range g.Thresholds {
		if value >= r.Value {
			return r.Icon
		}
	}
	return ""
}

// Threshold represents a Value-Icon pair
type Threshold struct {
	Value int
	Icon  string
}
