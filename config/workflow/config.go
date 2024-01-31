package workflow

import "github.com/crossplane/upjet/pkg/config"

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
				Type: "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/alert/v1alpha1.Channel",
			},
		}
	})
}
