package config

import (
	"okki-status/impl/provider"
	"reflect"
)

var providerTypeMap = map[string]reflect.Type{
	"battery": reflect.TypeOf(provider.Battery{}),
}
