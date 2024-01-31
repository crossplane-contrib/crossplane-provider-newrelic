package nrqldroprule

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_nrql_drop_rule", func(r *config.Resource) {
		r.ShortGroup = "nrql"
	})
}
