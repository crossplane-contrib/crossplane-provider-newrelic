package notification_destination

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_notification_destination", func(r *config.Resource) {
		r.ShortGroup = "alert"
		// https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/notification_destination#import
	})
}
