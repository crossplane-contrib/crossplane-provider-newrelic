package pipelinecloudrule

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pipeline_cloud_rule", func(r *config.Resource) {
		r.ShortGroup = "pipeline"
	})
}
