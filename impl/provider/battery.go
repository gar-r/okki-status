package provider

type Battery struct {
	Device string `yaml:"device"`
}

func (b *Battery) Status() string {
	return b.Device
}
