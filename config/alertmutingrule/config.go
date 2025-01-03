package alertmutingrule

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_alert_muting_rule", func(r *config.Resource) {
		r.ShortGroup = "alert"
		// https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/alert_muting_rule#import
	})
}
