/*
Copyright 2024 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/alert_policy"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboard"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/dashboard_json"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/log_parsing_rule"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notification_channel"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/notification_destination"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/nrql_alert_condition"
	"github.com/crossplane-contrib/crossplane-provider-newrelic/config/nrql_drop_rule"
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
		alert_policy.Configure,
		notification_channel.Configure,
		notification_destination.Configure,
		dashboard.Configure,
		dashboard_json.Configure,
		log_parsing_rule.Configure,
		nrql_alert_condition.Configure,
		nrql_drop_rule.Configure,
		workflow.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
