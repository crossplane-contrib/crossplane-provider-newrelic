/*
Copyright 2026 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	channel "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/alert/channel"
	destination "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/alert/destination"
	mutingrule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/alert/mutingrule"
	policy "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/alert/policy"
	dashboard "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/dashboard/dashboard"
	dashboardjson "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/dashboard/dashboardjson"
	parsingrule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/log/parsingrule"
	alertcondition "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/nrql/alertcondition"
	cloudrule "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/pipeline/cloudrule"
	providerconfig "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/providerconfig"
	workflow "github.com/crossplane-contrib/crossplane-provider-newrelic/internal/controller/cluster/workflow/workflow"
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
		cloudrule.Setup,
		providerconfig.Setup,
		workflow.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		channel.SetupGated,
		destination.SetupGated,
		mutingrule.SetupGated,
		policy.SetupGated,
		dashboard.SetupGated,
		dashboardjson.SetupGated,
		parsingrule.SetupGated,
		alertcondition.SetupGated,
		cloudrule.SetupGated,
		providerconfig.SetupGated,
		workflow.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
