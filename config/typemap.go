package config

import (
	"okki-status/provider"
	"reflect"
)

var TypeMap = map[string]reflect.Type{
	"clock":      reflect.TypeOf(provider.Clock{}),
	"battery":    reflect.TypeOf(provider.Battery{}),
	"layout":     reflect.TypeOf(provider.Layout{}),
	"volume":     reflect.TypeOf(provider.Volume{}),
	"memory":     reflect.TypeOf(provider.Memory{}),
	"brightness": reflect.TypeOf(provider.Brightness{}),
	"updates":    reflect.TypeOf(provider.Updates{}),
	"wireless":   reflect.TypeOf(provider.Wireless{}),
}
