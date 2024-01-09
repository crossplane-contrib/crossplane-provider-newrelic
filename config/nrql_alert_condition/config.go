package nrql_alert_condition

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_nrql_alert_condition", func(r *config.Resource) {
		r.ShortGroup = "nrql"

		// https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/nrql_alert_condition
		deprecatedItems := []string{"violation_time_limit", "nrql.evaluation_offset", "nrql.since_value", "term"}
		for _, deprecatedItem := range deprecatedItems {
			if s, ok := r.TerraformResource.Schema[deprecatedItem]; ok {
				s.Required = false
				s.Optional = false
			}
		}

		// Even though this is meant only for external name, the method above doesn't seem to work
		r.ExternalName.OmittedFields = append(r.ExternalName.OmittedFields, "term")

		r.References = config.References{
			"policy_id": config.Reference{
				Type: "github.com/crossplane-contrib/crossplane-provider-newrelic/apis/alert/v1alpha1.Policy",
			},
		}
	})
}
