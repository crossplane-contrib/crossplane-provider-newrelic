package workflow

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_workflow", func(r *config.Resource) {
		r.ShortGroup = "workflow"

		// Remove deprecated
		// if s, ok := r.TerraformResource.Schema["destinations_enabled"]; ok {
		//	s.Required = false
		//	s.Optional = false
		//}

		r.References = config.References{
			"destination.channel_id": config.Reference{
				TerraformName: "newrelic_notification_channel",
			},
		}
	})
}
