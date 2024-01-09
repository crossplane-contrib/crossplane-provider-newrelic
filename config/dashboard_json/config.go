package dashboard_json

import "github.com/crossplane/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("newrelic_one_dashboard_json", func(r *config.Resource) {
		r.ShortGroup = "dashboard"
		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		//r.References["dashboard"] = config.Reference{
		//	Type: "github.com/upbound/provider-newrelic/apis/dashboard/v1alpha1.Dashboard",
		//}
	})
}
