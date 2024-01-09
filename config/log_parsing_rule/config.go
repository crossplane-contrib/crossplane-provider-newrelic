package log_parsing_rule

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_log_parsing_rule", func(r *config.Resource) {
		r.ShortGroup = "log"
	})
}
