/*
Copyright 2026 Upbound Inc.
*/

//nolint:dupl // Similar structure to registry_namespaced.go but for cluster-scoped resources
package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/alertmutingrule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/alertpolicy"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboard"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboardjson"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/logparsingrule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notificationchannel"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notificationdestination"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/nrqlalertcondition"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/pipelinecloudrule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/workflow"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// GetProvider returns the provider configuration.
// The `generationProvider` argument specifies whether the provider
// configuration is being read for the code generation pipelines.
// In that case, we will only use the JSON schema for generating
// the CRDs.
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithRootGroup("newrelic.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		alertmutingrule.Configure,
		alertpolicy.Configure,
		notificationchannel.Configure,
		notificationdestination.Configure,
		dashboard.Configure,
		dashboardjson.Configure,
		logparsingrule.Configure,
		nrqlalertcondition.Configure,
		pipelinecloudrule.Configure,
		workflow.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
