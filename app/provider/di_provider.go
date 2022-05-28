package provider

import "github.com/fgrosse/goldi"

func registerTypes(types goldi.TypeRegistry) {
	types.RegisterAll(map[string]goldi.TypeFactory{})
}
