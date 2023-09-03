package config

import (
	"errors"
	"fmt"
	"okki-status/core"
	"reflect"

	"gopkg.in/yaml.v3"
)

func initProvider(m *core.Module) error {
	tname, err := typename(m.ProviderConf)
	if err != nil {
		return err
	}
	ptype, ok := providerTypeMap[tname]
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

const (
	errTypeMissing     = "missing type attribute"
	errUnknownProvider = "unknown provider: %s"
)
