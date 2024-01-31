package notificationchannel

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_notification_channel", func(r *config.Resource) {
		r.ShortGroup = "alert"
		// https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/notification_channel#import
		r.References = config.References{
			"destination_id": config.Reference{
				Type: "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/alert/v1alpha1.Destination",
			},
		}
	})
}
