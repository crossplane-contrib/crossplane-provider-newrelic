/*
Copyright 2024 Upbound Inc.
*/

package config

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

// TODO: Inject account_id on Channel, Destination, alert condition, drop rule
// NR - please add ID in destination in the UI

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"newrelic_alert_policy":             configAlertPolicy(),
	"newrelic_alert_muting_rule":        config.IdentifierFromProvider,
	"newrelic_notification_channel":     config.IdentifierFromProvider,
	"newrelic_notification_destination": config.IdentifierFromProvider,
	"newrelic_one_dashboard":            configDashboard(),
	"newrelic_one_dashboard_json":       config.IdentifierFromProvider,
	"newrelic_log_parsing_rule":         config.IdentifierFromProvider,
	"newrelic_nrql_alert_condition":     configNrqAlertCondition(),
	"newrelic_nrql_drop_rule":           configNrqDropRule(),
	"newrelic_workflow":                 config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

// See - https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/alert_policy#import
// Alert policies can be imported using a composite ID of <id>:<account_id>,
// where account_id is the account number scoped to the alert policy resource.
func configAlertPolicy() config.ExternalName {
	e := config.IdentifierFromProvider

	// Gets the policy id as the external-name even though terraform internal uses policy_id:account_id
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("unable to get id from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("Can't format id from tfstate as string")
		}
		// <id>:<account_id>
		words := strings.Split(idStr, ":")
		return words[0], nil
	}

	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		// Using a stub value to pass validation.
		if len(externalName) == 0 {
			return "", nil
		}

		accountID, ok := parameters["account_id"]
		if !ok {
			return "", errors.New("Can't get account_id from alert policy")
		}
		accountIDStr := strconv.FormatFloat(accountID.(float64), 'f', 0, 64)

		// <id>:<account_id>
		return fmt.Sprintf("%s:%s", externalName, accountIDStr), nil
	}
	return e
}

// See - https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/nrql_alert_condition#import
// NRQL alert conditions can be imported using a composite ID of <policy_id>:<condition_id>, e.g.
func configNrqAlertCondition() config.ExternalName {
	e := config.IdentifierFromProvider

	// Gets the condition id as the external-name even though terraform internal uses <policy_id>:<condition_id>:<conditionType>
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("unable to get id from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("Can't format id from tfstate as string")
		}
		// <policy_id>:<condition_id>
		words := strings.Split(idStr, ":")
		return words[1], nil
	}

	// See - https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/nrql_alert_condition#import
	// Formats the condition id as <policy_id>:<condition_id>
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		// Using a stub value to pass validation.
		if len(externalName) == 0 {
			return "", nil
		}

		policyID, ok := parameters["policy_id"]
		if !ok {
			return "", errors.New("Can't get policy_id from nrql alert condition")
		}
		policyIDStr := strconv.FormatFloat(policyID.(float64), 'f', 0, 64)

		// <policy_id>:<condition_id>
		return fmt.Sprintf("%s:%s", policyIDStr, externalName), nil
	}
	return e
}

// See - https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/one_dashboard#import
func configDashboard() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("unable to get id from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("Can't format id from tfstate as string")
		}
		return idStr, nil
	}
	return e
}

// See - https://registry.terraform.io/providers/newrelic/newrelic/latest/docs/resources/nrql_drop_rule#import
// New Relic NRQL drop rules can be imported using a concatenated string of the format <account_id>:<rule_id>
func configNrqDropRule() config.ExternalName {
	e := config.IdentifierFromProvider

	// Gets the rule id as the external-name even though terraform internal uses account_id:rule_id
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("unable to get id from tfstate")
		}
		idStr, ok := id.(string)
		if !ok {
			return "", errors.New("Can't format id from tfstate as string")
		}
		// <id>:<account_id>
		words := strings.Split(idStr, ":")
		return words[1], nil
	}

	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		// Using a stub value to pass validation.
		if len(externalName) == 0 {
			return "", nil
		}

		accountID, ok := parameters["account_id"]
		if !ok {
			return "", errors.New("Can't get account_id from nrql_drop_rule")
		}
		accountIDStr := strconv.FormatFloat(accountID.(float64), 'f', 0, 64)

		// <account_id>:<rule_id>
		return fmt.Sprintf("%s:%s", accountIDStr, externalName), nil
	}
	return e
}
