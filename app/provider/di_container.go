package provider

import "github.com/fgrosse/goldi"

var _ *goldi.Container

func init() {
	registry := goldi.NewTypeRegistry()
	registerTypes(registry)
	_ = goldi.NewContainer(registry, nil)
}
