/*
Copyright 2024 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/alertpolicy"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboard"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboardjson"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/logparsingrule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notificationchannel"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notificationdestination"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/nrqlalertcondition"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/nrqldroprule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/workflow"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "newrelic"
	modulePath     = "github.com/crossplane-contrib/crossplane-provider-newrelic"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("newrelic.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		alertpolicy.Configure,
		notificationchannel.Configure,
		notificationdestination.Configure,
		dashboard.Configure,
		dashboardjson.Configure,
		logparsingrule.Configure,
		nrqlalertcondition.Configure,
		nrqldroprule.Configure,
		workflow.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
