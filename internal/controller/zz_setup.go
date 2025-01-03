/*
Copyright 2024 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	channel "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/alert/channel"
	destination "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/alert/destination"
	mutingrule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/alert/mutingrule"
	policy "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/alert/policy"
	dashboard "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/dashboard/dashboard"
	dashboardjson "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/dashboard/dashboardjson"
	parsingrule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/log/parsingrule"
	alertcondition "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/nrql/alertcondition"
	droprule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/nrql/droprule"
	providerconfig "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/providerconfig"
	workflow "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/workflow/workflow"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.Setup,
		destination.Setup,
		mutingrule.Setup,
		policy.Setup,
		dashboard.Setup,
		dashboardjson.Setup,
		parsingrule.Setup,
		alertcondition.Setup,
		droprule.Setup,
		providerconfig.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
