# Provider NewRelic

`crossplane-provider-newrelic` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
NewRelic API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/crossplane-contrib/crossplane-provider-newrelic):
```
up ctp provider install crossplane-contrib/crossplane-provider-newrelic:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-newrelic
spec:
  package: crossplane-contrib/crossplane-provider-newrelic:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/upbound/provider-newrelic).

## Authentication
This is documented in [AUTHENTICATION.md](AUTHENTICATION.md).
* [examples/providerconfig/providerconfig.yaml](examples/providerconfig/providerconfig.yaml)
* [examples/providerconfig/secret.yaml.tmpl](examples/providerconfig/secret.yaml.tmpl)

# Currently Supported Resources
* `alert_channel` - [examples-generated/alert/channel.yaml](examples-generated/alert/channel.yaml).
* `alert_destination` - [examples/alert/destination.yaml](examples/alert/destination.yaml).
* `alert_muting_rule` - [examples-generated/alert/mutingrule.yaml](examples-generated/alert/channel.yaml).
* `alert_policy` - [examples/alert/policy.yaml](examples/alert/policy.yaml).
* `dashboardjson` - [examples/dashboard/dashboardjson.yaml](examples/dashboard/dashboardjson.yaml).
* `nrql_alert_condition` - [examples-generated/nrql/alertcondition.yaml](examples-generated/nrql/alertcondition.yaml).
* `workflow` - [examples/workflow/workflow.yaml](examples/workflow/workflow.yaml).

## Upstream Documentation
* https://github.com/newrelic/terraform-provider-newrelic/tree/main/newrelic
* https://github.com/upbound/upjet-provider-template
* https://github.com/crossplane/upjet/blob/v0.10.0/docs/add-new-resource-short.md

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contribb/crossplane-provider-newrelic/issues).
