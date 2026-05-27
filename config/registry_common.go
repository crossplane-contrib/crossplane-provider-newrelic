/*
Copyright 2026 Upbound Inc.
*/

package config

import (
	_ "embed"
)

var (
	//go:embed schema.json
	providerSchema string

	//go:embed provider-metadata.yaml
	providerMetadata []byte
)

const (
	resourcePrefix = "newrelic"
	modulePath     = "github.com/crossplane-contrib/crossplane-provider-newrelic"
)
