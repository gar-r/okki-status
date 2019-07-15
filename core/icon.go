package core

type StaticIcon struct {
	Icon string
}

func (i *StaticIcon) GetIcon(status string) string {
	return i.Icon
}

type ThresholdIcon struct {
	Thresholds        []Threshold
	StatusConverterFn func(string) int
}

func (g *ThresholdIcon) GetIcon(status string) string {
	value := g.StatusConverterFn(status)
	for _, r := range g.Thresholds {
		if value >= r.Value {
			return r.Icon
		}
	}
	return ""
}

type Threshold struct {
	Value int
	Icon  string
}
