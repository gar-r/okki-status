package config

import (
	"okki-status/provider"
	"reflect"
)

var TypeMap = map[string]reflect.Type{
	"battery": reflect.TypeOf(provider.Battery{}),
}
