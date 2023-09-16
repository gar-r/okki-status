package config

import (
	"errors"
	"fmt"
	"io"
	"okki-status/core"
	"reflect"

	"gopkg.in/yaml.v3"
)

const (
	errTypeMissing     = "missing type attribute"
	errUnknownProvider = "unknown provider: %s"
)

func Parse(r io.Reader) (*core.Bar, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	bar := &core.Bar{}
	err = yaml.Unmarshal(data, bar)
	if err != nil {
		return nil, err
	}
	err = initBar(bar)
	return bar, err
}

func initBar(b *core.Bar) error {
	for _, m := range b.Modules {
		err := initModule(m)
		if err != nil {
			return fmt.Errorf("module %s cannot be initialized: %s", m.Name, err)
		}
	}
	return nil
}

func initModule(m *core.Module) error {
	tname, err := typename(m.ProviderConf)
	if err != nil {
		return err
	}
	// when missing, default module name to provider type name
	if m.Name == "" {
		m.Name = tname
	}
	// pre-compile variant regular expressions
	for i, v := range m.Variants {
		if err := v.Compile(); err != nil {
			return fmt.Errorf("module %s variant #%d contains an invalid Pattern: %s, %s", m.Name, i, v.Pattern, err)
		}
	}

	ptype, ok := TypeMap[tname]
	if !ok {
		return fmt.Errorf(errUnknownProvider, tname)
	}
	m.Provider = reflect.New(ptype).Interface().(core.Provider)
	b, err := yaml.Marshal(m.ProviderConf)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, m.Provider)
}

func typename(pconf map[string]interface{}) (string, error) {
	tstr, ok := pconf["type"]
	if !ok {
		return "", errors.New(errTypeMissing)
	}
	return tstr.(string), nil
}
