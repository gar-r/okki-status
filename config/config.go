package config

import (
	"fmt"
	"io"
	"okki-status/core"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Modules []*core.Module `yaml:"modules"`
}

func Read(r io.Reader) (*Conf, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	conf := &Conf{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	err = conf.initialize()
	return conf, err
}

func (c *Conf) initialize() error {
	for _, m := range c.Modules {
		err := initProvider(m)
		if err != nil {
			return fmt.Errorf("module initialization error: %s", m.Name)
		}
	}
	return nil
}
