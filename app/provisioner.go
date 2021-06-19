package app

import "github.com/cosmos/cosmos-sdk/container"

type Provisioner interface {
	Provision(key ModuleKey, registrar container.Registrar) error
}

type Provider interface {
	Provide(key ModuleKey) ([]interface{}, error)
}